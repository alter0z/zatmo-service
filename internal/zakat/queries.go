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

const GetVillager = `
	SELECT 
		v.id,
		v.villager
	FROM villager AS v
	ORDER BY created_at ASC`

const CreateZakat = `
	INSERT INTO zakat (villager, name, total_people, amount, charity, category)
	VALUES (
		$1,
		$2,
		$3,
		CASE 
			WHEN $6 = TRUE THEN $4::numeric * 3.25
			ELSE $4::numeric * 40000
		END,
		$5,
		$6
	)
	RETURNING id, created_at, updated_at`

const UpdateZakat = `
	UPDATE zakat
	SET
		name = $2,
		total_people = $3,
		amount = CASE
								WHEN $6 = TRUE THEN $4::numeric * 3.25
								ELSE $4::numeric * 40000
							END,
		charity = $5,
		category = $6,
		updated_at = NOW()
	WHERE id = $1
	RETURNING created_at, updated_at`

const DeleteZakat = `DELETE FROM zakat WHERE id = $1`