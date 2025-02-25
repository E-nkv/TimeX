--today
select * from sessions WHERE 
    start > CURRENT_DATE - '1 second' :: INTERVAL AND 
    start < CURRENT_DATE + '23 hours 59 minutes'::INTERVAL;

--any day: givenDay = '2025-02-12'
select * from sessions WHERE 
    start > '2025-02-12'::DATE - '1 second' :: INTERVAL AND 
    start < '2025-02-12'::DATE + '23 hours 59 minutes'::INTERVAL;


--this week
SELECT * FROM sessions WHERE 
    start >= CURRENT_DATE - ((COALESCE(NULLIF(EXTRACT(DOW FROM CURRENT_DATE), 0), 7)) || ' days')::INTERVAL +  '23 hours 59 minutes'::INTERVAL AND 
    start <= CURRENT_DATE + ((7 - (COALESCE(NULLIF(EXTRACT(DOW FROM CURRENT_DATE), 0), 7))) || ' days')::INTERVAL + '23 hours 59 minutes'::INTERVAL;


--any week, given the date of a day of that week '2025-02-12'
SELECT * FROM sessions WHERE 
    start >= '2025-02-12'::DATE - ((COALESCE(NULLIF(EXTRACT(DOW FROM '2025-02-12'::DATE), 0), 7)) || ' days')::INTERVAL AND 
    start <= '2025-02-12'::DATE + ((7 - (COALESCE(NULLIF(EXTRACT(DOW FROM '2025-02-12'::DATE), 0), 7))) || ' days')::INTERVAL + '23 hours 59 minutes'::INTERVAL;


--this month
SELECT * FROM sessions WHERE start >= DATE_TRUNC('month', CURRENT_DATE) - '1 day' :: INTERVAL AND "end" <= DATE_TRUNC('month', CURRENT_DATE + '1 month'::INTERVAL) - '1 day' ::INTERVAL;

--any month, given the date of a day of that month '2025-02-12'
SELECT * FROM sessions WHERE 
    start > DATE_TRUNC('month', DATE'2025-02-12') - '1 day' :: INTERVAL + '23 hours 59 minutes'::INTERVAL AND 
    "end" < DATE_TRUNC('month', DATE'2025-02-12' + '1 month'::INTERVAL) - '1 day' ::INTERVAL;

--this year
SELECT * FROM sessions WHERE
    start > DATE_TRUNC(CURRENT_DATE, 'year') - '1 day 1 second'::INTERVAL AND
    "end" < DATE_TRUNC(CURRENT_DATE + '1 year' :: INTERVAL, 'year') - '1 day 1 second' ::INTERVAL

--any year, given the year = '2024'
SELECT * FROM sessions WHERE
    start >  MAKE_TIMESTAMPTZ(2024, 1, 1, 0, 0, 0) - '1 day 1 second'::INTERVAL AND
    "end" < MAKE_TIMESTAMPTZ(2024 + 1, 1, 1, 0, 0, 0) - '1 day 1 second' ::INTERVAL


--graph, timemode = day
select (id, start, "end", category_id, focus, delta, category_line) from sessions WHERE 
    start > $1::DATE - '1 second' :: INTERVAL AND 
    start < $1::DATE + '23 hours 59 minutes'::INTERVAL AND
    ($2 = -1 OR $2 = ANY(category_line))
    ORDER BY start;

--graph, timemode = month. Input: $1=Date (with Month-precision), $2=category_id
WITH 
    DaySeries AS (
        SELECT GENERATE_SERIES(1, (SELECT EXTRACT(DAY FROM (date_trunc('month', $1::DATE) + INTERVAL '1 month - 1 day')))) AS day
    ),
    SessionTotals AS (
        SELECT EXTRACT(DAY FROM DATE_TRUNC('day', start)) AS day, SUM(delta) AS tot 
        FROM sessions 
        WHERE  
            start >= DATE_TRUNC('month', $1::DATE) - INTERVAL '1 day' AND 
            "end" <= DATE_TRUNC('month', $1::DATE + INTERVAL '1 month') - INTERVAL '1 day' AND
            ($2 = -1 OR $2 = ANY(category_line))
        GROUP BY day
    ),
    SuperTotal AS (
        SELECT SUM(tot) AS super_tot FROM SessionTotals
    )
SELECT 
    NULL AS day, 
    COALESCE(st.super_tot, INTERVAL '0 seconds') AS tot
FROM SuperTotal st
UNION ALL
SELECT 
    ds.day, 
    COALESCE(st.tot, INTERVAL '0 seconds') AS tot
FROM DaySeries ds
LEFT JOIN SessionTotals st ON ds.day = st.day
ORDER BY day NULLS FIRST;

--graph, timemode = year. $1 = year, $2 = category_id
WITH MonthSeries AS (
    SELECT generate_series(1, 12) AS month
),
SessionTotals AS (
    SELECT 
        EXTRACT(MONTH FROM DATE_TRUNC('month', start)) AS month, 
        SUM(delta) AS tot 
    FROM sessions 
    WHERE 
        start > MAKE_TIMESTAMPTZ((EXTRACT (YEAR FROM $1::DATE) :: INT ), 1, 1, 0, 0, 0) - INTERVAL '1 day 1 second' AND
        "end" < MAKE_TIMESTAMPTZ((EXTRACT (YEAR FROM $1::DATE) :: INT + 1), 1, 1, 0, 0, 0) - INTERVAL '1 day 1 second' AND
        ($2 = -1 OR $2 = ANY(category_line))
    GROUP BY month
),
SuperTotal AS (
    SELECT SUM(tot) AS super_tot FROM SessionTotals
)
SELECT 
    NULL AS month, 
    COALESCE(st.super_tot, INTERVAL '0 seconds') AS tot
FROM SuperTotal st

UNION ALL

SELECT 
    ms.month, 
    COALESCE(st.tot, INTERVAL '0 seconds') AS tot
FROM MonthSeries ms
LEFT JOIN SessionTotals st ON ms.month = st.month
ORDER BY month NULLS FIRST;

--pie, timemode = day ($1, in '2024-02-15'), category = $2

select (id, start, "end", category_id, focus, delta, category_line) from sessions WHERE 
    start > $1::DATE - '1 second' :: INTERVAL AND 
    start < $2::DATE + '23 hours 59 minutes'::INTERVAL AND
    $2 = ANY(category_line);
