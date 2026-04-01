package zakat

const GetZakat = `
	SELECT 
		z.id, 
		z.created_at, 
		z.updated_at,
		v.villager, 
		z.name, 
		z.total_people, 
		z.amount, 
		z.charity , 
		z.category
	FROM zakat as z
	JOIN villager AS v ON v.id = z.villager
	WHERE
		($1::uuid IS NULL OR z.villager = $1::uuid)
		AND ($2::text IS NULL OR z.name ILIKE '%' || $2::text || '%')
		AND ($3::boolean IS NULL OR z.category = $3::boolean)
		AND ($4::date IS NULL OR (z.created_at AT TIME ZONE 'Asia/Jakarta')::date = $4::date)
	ORDER BY created_at desc`