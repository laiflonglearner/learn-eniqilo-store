package postgres

import (
	"fmt"
	"strings"

	"github.com/backend-magang/eniqilo-store/models/entity"
	"github.com/spf13/cast"
)

func buildQueryGetListProducts(req entity.GetListProductRequest) (string, []interface{}) {
	var (
		queryBuilder  strings.Builder
		args          []interface{}
		orderByClause string
	)

	queryBuilder.WriteString("SELECT * FROM products WHERE 1=1")

	if req.ID != "" {
		queryBuilder.WriteString(" AND id = ?")
		args = append(args, req.ID)
	}

	if req.Name != "" {
		queryBuilder.WriteString(" AND name ILIKE ?")
		args = append(args, "%"+req.Name+"%")
	}

	if req.IsAvailable != "" {
		queryBuilder.WriteString(" AND is_available = ?")
		args = append(args, cast.ToBool(req.IsAvailable))
	}

	if req.InStock != "" {
		inStock := cast.ToBool(req.InStock)
		if inStock {
			queryBuilder.WriteString(" AND stock > 0")
		} else {
			queryBuilder.WriteString(" AND stock <= 0")
		}
	}

	if req.Category != "" {
		queryBuilder.WriteString(" AND category = ?")
		args = append(args, req.Category)
	}

	if req.SKU != "" {
		queryBuilder.WriteString(" AND sku = ?")
		args = append(args, req.SKU)
	}

	queryBuilder.WriteString(" AND deleted_at IS NULL")

	if req.Price != "" {
		orderByClause = fmt.Sprintf("price %s", strings.ToUpper(req.Price))
	}

	if req.CreatedAt != "" {
		if orderByClause != "" {
			orderByClause += fmt.Sprintf(", created_at %s", strings.ToUpper(req.CreatedAt))
		} else {
			orderByClause += fmt.Sprintf("created_at %s", strings.ToUpper(req.CreatedAt))
		}
	}

	if orderByClause != "" {
		queryBuilder.WriteString(" ORDER BY " + orderByClause)
	} else {
		queryBuilder.WriteString(" ORDER BY created_at DESC")
	}

	queryBuilder.WriteString(" LIMIT ? OFFSET ?")
	args = append(args, cast.ToInt(req.Limit), cast.ToInt(req.Offset))

	return queryBuilder.String(), args
}
