-- Encoding name
DROP FUNCTION IF EXISTS test_encoding_name(int);
CREATE OR REPLACE FUNCTION test_encoding_name(int) RETURNS text AS '$libdir/gppc_test' LANGUAGE c IMMUTABLE;
SELECT
       test_encoding_name(0),
       test_encoding_name(1),
       test_encoding_name(2),
       test_encoding_name(3),
       test_encoding_name(4),
       test_encoding_name(5),
       test_encoding_name(6),
       test_encoding_name(7),
       test_encoding_name(8),
       test_encoding_name(9),
       test_encoding_name(10),
       test_encoding_name(11),
       test_encoding_name(12),
       test_encoding_name(13),
       test_encoding_name(14),
       test_encoding_name(15),
       test_encoding_name(16),
       test_encoding_name(17),
       test_encoding_name(18),
       test_encoding_name(19),
       test_encoding_name(20),
       test_encoding_name(21),
       test_encoding_name(22),
       test_encoding_name(23),
       test_encoding_name(24),
       test_encoding_name(25),
       test_encoding_name(26),
       test_encoding_name(27),
       test_encoding_name(28),
       test_encoding_name(29),
       test_encoding_name(30),
       test_encoding_name(31),
       test_encoding_name(32),
       test_encoding_name(33),
       test_encoding_name(34);