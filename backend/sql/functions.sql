CREATE OR REPLACE FUNCTION category_line(input_category_id INT)
RETURNS INT[] AS $$
BEGIN
  RETURN (
    WITH RECURSIVE category_path AS (
      SELECT id, parent_category_id, ARRAY[id] AS path
      FROM categories
      WHERE id = input_category_id

      UNION ALL

      SELECT c.id, c.parent_category_id, cp.path || c.id
      FROM categories c
      JOIN category_path cp ON c.id = cp.parent_category_id
    )
    SELECT array_reverse(path)
    FROM category_path
    WHERE parent_category_id IS NULL
  );
END;
$$
LANGUAGE plpgsql IMMUTABLE;



CREATE OR REPLACE FUNCTION array_reverse(anyarray)
RETURNS anyarray AS $$
SELECT ARRAY(SELECT $1[i] FROM generate_subscripts($1,1) AS s(i) ORDER BY i DESC);
$$ LANGUAGE 'sql' STRICT IMMUTABLE;
