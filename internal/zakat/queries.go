package zakat

const GetZakatAll = `
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
	ORDER BY created_at desc`