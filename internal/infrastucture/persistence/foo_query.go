package persistence

var (
	fooQueries = struct {
		From string
		All  string
	}{
		From: `FROM foo f `,
		All: `
			f.id,
			f.receipt_name,
			f.receipt_notes,
			f.total_price,
			f.is_active,
			f.created_time,
			f.created_by,
			f.updated_time,
			f.updated_by,
			f.removed_time,
			f.removed_by
		`,
	}
)
