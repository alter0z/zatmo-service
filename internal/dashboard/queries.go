package dashboard

const GetSummaryCounts = `
  SELECT 
    COALESCE(SUM(z.total_people), 0) AS total_zakat,
    COALESCE(SUM(CASE WHEN z.category = true  THEN z.total_people ELSE 0 END), 0) AS total_rice_zakat,
    COALESCE(SUM(CASE WHEN z.category = false THEN z.total_people ELSE 0 END), 0) AS total_cash_zakat,
    COALESCE(SUM(CASE WHEN z.category = true  THEN z.charity     ELSE 0 END), 0) AS total_rice_charity,
    COALESCE(SUM(CASE WHEN z.category = false THEN z.charity     ELSE 0 END), 0) AS total_cash_charity,
    COALESCE(SUM(CASE WHEN z.category = true  THEN z.amount      ELSE 0 END), 0) AS total_rice_amount,
    COALESCE(SUM(CASE WHEN z.category = false THEN z.amount      ELSE 0 END), 0) AS total_cash_amount
  FROM zakat as z
  WHERE ($1::uuid IS NULL OR z.villager = $1::uuid);`

const GetReceiverCounts = `
  SELECT 
    (SELECT COALESCE(SUM(total), 0) FROM needy)      AS total_needy,
    (SELECT COALESCE(SUM(total), 0) FROM destitute)  AS total_destitute;`