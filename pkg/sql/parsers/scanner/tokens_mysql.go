// Code generated by goyacc -o mysql_sql.go -c mysql mysql_sql.y. DO NOT EDIT.
package scanner

const MYSQL_LEX_ERROR = 57346
const MYSQL_EMPTY = 57347
const MYSQL_UNION = 57348
const MYSQL_SELECT = 57349
const MYSQL_STREAM = 57350
const MYSQL_INSERT = 57351
const MYSQL_UPDATE = 57352
const MYSQL_DELETE = 57353
const MYSQL_FROM = 57354
const MYSQL_WHERE = 57355
const MYSQL_GROUP = 57356
const MYSQL_HAVING = 57357
const MYSQL_ORDER = 57358
const MYSQL_BY = 57359
const MYSQL_LIMIT = 57360
const MYSQL_OFFSET = 57361
const MYSQL_FOR = 57362
const MYSQL_LOWER_THAN_SET = 57363
const MYSQL_SET = 57364
const MYSQL_ALL = 57365
const MYSQL_DISTINCT = 57366
const MYSQL_DISTINCTROW = 57367
const MYSQL_AS = 57368
const MYSQL_EXISTS = 57369
const MYSQL_ASC = 57370
const MYSQL_DESC = 57371
const MYSQL_INTO = 57372
const MYSQL_DUPLICATE = 57373
const MYSQL_DEFAULT = 57374
const MYSQL_LOCK = 57375
const MYSQL_KEYS = 57376
const MYSQL_VALUES = 57377
const MYSQL_NEXT = 57378
const MYSQL_VALUE = 57379
const MYSQL_SHARE = 57380
const MYSQL_MODE = 57381
const MYSQL_SQL_NO_CACHE = 57382
const MYSQL_SQL_CACHE = 57383
const MYSQL_JOIN = 57384
const MYSQL_STRAIGHT_JOIN = 57385
const MYSQL_LEFT = 57386
const MYSQL_RIGHT = 57387
const MYSQL_INNER = 57388
const MYSQL_OUTER = 57389
const MYSQL_CROSS = 57390
const MYSQL_NATURAL = 57391
const MYSQL_USE = 57392
const MYSQL_FORCE = 57393
const MYSQL_ON = 57394
const MYSQL_USING = 57395
const MYSQL_SUBQUERY_AS_EXPR = 57396
const MYSQL_LOWER_THAN_STRING = 57397
const MYSQL_ID = 57398
const MYSQL_AT_ID = 57399
const MYSQL_AT_AT_ID = 57400
const MYSQL_STRING = 57401
const MYSQL_VALUE_ARG = 57402
const MYSQL_LIST_ARG = 57403
const MYSQL_COMMENT = 57404
const MYSQL_COMMENT_KEYWORD = 57405
const MYSQL_INTEGRAL = 57406
const MYSQL_HEX = 57407
const MYSQL_BIT_LITERAL = 57408
const MYSQL_FLOAT = 57409
const MYSQL_HEXNUM = 57410
const MYSQL_NULL = 57411
const MYSQL_TRUE = 57412
const MYSQL_FALSE = 57413
const MYSQL_LOWER_THAN_CHARSET = 57414
const MYSQL_CHARSET = 57415
const MYSQL_UNIQUE = 57416
const MYSQL_KEY = 57417
const MYSQL_OR = 57418
const MYSQL_PIPE_CONCAT = 57419
const MYSQL_XOR = 57420
const MYSQL_AND = 57421
const MYSQL_NOT = 57422
const MYSQL_BETWEEN = 57423
const MYSQL_CASE = 57424
const MYSQL_WHEN = 57425
const MYSQL_THEN = 57426
const MYSQL_ELSE = 57427
const MYSQL_END = 57428
const MYSQL_LE = 57429
const MYSQL_GE = 57430
const MYSQL_NE = 57431
const MYSQL_NULL_SAFE_EQUAL = 57432
const MYSQL_IS = 57433
const MYSQL_LIKE = 57434
const MYSQL_REGEXP = 57435
const MYSQL_IN = 57436
const MYSQL_ASSIGNMENT = 57437
const MYSQL_SHIFT_LEFT = 57438
const MYSQL_SHIFT_RIGHT = 57439
const MYSQL_DIV = 57440
const MYSQL_MOD = 57441
const MYSQL_UNARY = 57442
const MYSQL_COLLATE = 57443
const MYSQL_BINARY = 57444
const MYSQL_UNDERSCORE_BINARY = 57445
const MYSQL_INTERVAL = 57446
const MYSQL_BEGIN = 57447
const MYSQL_START = 57448
const MYSQL_TRANSACTION = 57449
const MYSQL_COMMIT = 57450
const MYSQL_ROLLBACK = 57451
const MYSQL_WORK = 57452
const MYSQL_CONSISTENT = 57453
const MYSQL_SNAPSHOT = 57454
const MYSQL_CHAIN = 57455
const MYSQL_NO = 57456
const MYSQL_RELEASE = 57457
const MYSQL_PRIORITY = 57458
const MYSQL_QUICK = 57459
const MYSQL_BIT = 57460
const MYSQL_TINYINT = 57461
const MYSQL_SMALLINT = 57462
const MYSQL_MEDIUMINT = 57463
const MYSQL_INT = 57464
const MYSQL_INTEGER = 57465
const MYSQL_BIGINT = 57466
const MYSQL_INTNUM = 57467
const MYSQL_REAL = 57468
const MYSQL_DOUBLE = 57469
const MYSQL_FLOAT_TYPE = 57470
const MYSQL_DECIMAL = 57471
const MYSQL_NUMERIC = 57472
const MYSQL_DECIMAL_VALUE = 57473
const MYSQL_TIME = 57474
const MYSQL_TIMESTAMP = 57475
const MYSQL_DATETIME = 57476
const MYSQL_YEAR = 57477
const MYSQL_CHAR = 57478
const MYSQL_VARCHAR = 57479
const MYSQL_BOOL = 57480
const MYSQL_CHARACTER = 57481
const MYSQL_VARBINARY = 57482
const MYSQL_NCHAR = 57483
const MYSQL_TEXT = 57484
const MYSQL_TINYTEXT = 57485
const MYSQL_MEDIUMTEXT = 57486
const MYSQL_LONGTEXT = 57487
const MYSQL_BLOB = 57488
const MYSQL_TINYBLOB = 57489
const MYSQL_MEDIUMBLOB = 57490
const MYSQL_LONGBLOB = 57491
const MYSQL_JSON = 57492
const MYSQL_ENUM = 57493
const MYSQL_GEOMETRY = 57494
const MYSQL_POINT = 57495
const MYSQL_LINESTRING = 57496
const MYSQL_POLYGON = 57497
const MYSQL_GEOMETRYCOLLECTION = 57498
const MYSQL_MULTIPOINT = 57499
const MYSQL_MULTILINESTRING = 57500
const MYSQL_MULTIPOLYGON = 57501
const MYSQL_INT1 = 57502
const MYSQL_INT2 = 57503
const MYSQL_INT3 = 57504
const MYSQL_INT4 = 57505
const MYSQL_INT8 = 57506
const MYSQL_SQL_SMALL_RESULT = 57507
const MYSQL_SQL_BIG_RESULT = 57508
const MYSQL_SQL_BUFFER_RESULT = 57509
const MYSQL_LOW_PRIORITY = 57510
const MYSQL_HIGH_PRIORITY = 57511
const MYSQL_DELAYED = 57512
const MYSQL_CREATE = 57513
const MYSQL_ALTER = 57514
const MYSQL_DROP = 57515
const MYSQL_RENAME = 57516
const MYSQL_ANALYZE = 57517
const MYSQL_ADD = 57518
const MYSQL_SCHEMA = 57519
const MYSQL_TABLE = 57520
const MYSQL_INDEX = 57521
const MYSQL_VIEW = 57522
const MYSQL_TO = 57523
const MYSQL_IGNORE = 57524
const MYSQL_IF = 57525
const MYSQL_PRIMARY = 57526
const MYSQL_COLUMN = 57527
const MYSQL_CONSTRAINT = 57528
const MYSQL_SPATIAL = 57529
const MYSQL_FULLTEXT = 57530
const MYSQL_FOREIGN = 57531
const MYSQL_KEY_BLOCK_SIZE = 57532
const MYSQL_SHOW = 57533
const MYSQL_DESCRIBE = 57534
const MYSQL_EXPLAIN = 57535
const MYSQL_DATE = 57536
const MYSQL_ESCAPE = 57537
const MYSQL_REPAIR = 57538
const MYSQL_OPTIMIZE = 57539
const MYSQL_TRUNCATE = 57540
const MYSQL_MAXVALUE = 57541
const MYSQL_PARTITION = 57542
const MYSQL_REORGANIZE = 57543
const MYSQL_LESS = 57544
const MYSQL_THAN = 57545
const MYSQL_PROCEDURE = 57546
const MYSQL_TRIGGER = 57547
const MYSQL_STATUS = 57548
const MYSQL_VARIABLES = 57549
const MYSQL_ROLE = 57550
const MYSQL_PROXY = 57551
const MYSQL_AVG_ROW_LENGTH = 57552
const MYSQL_STORAGE = 57553
const MYSQL_DISK = 57554
const MYSQL_MEMORY = 57555
const MYSQL_CHECKSUM = 57556
const MYSQL_COMPRESSION = 57557
const MYSQL_DATA = 57558
const MYSQL_DIRECTORY = 57559
const MYSQL_DELAY_KEY_WRITE = 57560
const MYSQL_ENCRYPTION = 57561
const MYSQL_ENGINE = 57562
const MYSQL_MAX_ROWS = 57563
const MYSQL_MIN_ROWS = 57564
const MYSQL_PACK_KEYS = 57565
const MYSQL_ROW_FORMAT = 57566
const MYSQL_STATS_AUTO_RECALC = 57567
const MYSQL_STATS_PERSISTENT = 57568
const MYSQL_STATS_SAMPLE_PAGES = 57569
const MYSQL_DYNAMIC = 57570
const MYSQL_COMPRESSED = 57571
const MYSQL_REDUNDANT = 57572
const MYSQL_COMPACT = 57573
const MYSQL_FIXED = 57574
const MYSQL_COLUMN_FORMAT = 57575
const MYSQL_AUTO_RANDOM = 57576
const MYSQL_RESTRICT = 57577
const MYSQL_CASCADE = 57578
const MYSQL_ACTION = 57579
const MYSQL_PARTIAL = 57580
const MYSQL_SIMPLE = 57581
const MYSQL_CHECK = 57582
const MYSQL_ENFORCED = 57583
const MYSQL_RANGE = 57584
const MYSQL_LIST = 57585
const MYSQL_ALGORITHM = 57586
const MYSQL_LINEAR = 57587
const MYSQL_PARTITIONS = 57588
const MYSQL_SUBPARTITION = 57589
const MYSQL_SUBPARTITIONS = 57590
const MYSQL_TYPE = 57591
const MYSQL_ANY = 57592
const MYSQL_SOME = 57593
const MYSQL_PREPARE = 57594
const MYSQL_DEALLOCATE = 57595
const MYSQL_PROPERTIES = 57596
const MYSQL_PARSER = 57597
const MYSQL_VISIBLE = 57598
const MYSQL_INVISIBLE = 57599
const MYSQL_BTREE = 57600
const MYSQL_HASH = 57601
const MYSQL_RTREE = 57602
const MYSQL_BSI = 57603
const MYSQL_ZONEMAP = 57604
const MYSQL_LEADING = 57605
const MYSQL_BOTH = 57606
const MYSQL_TRAILING = 57607
const MYSQL_UNKNOWN = 57608
const MYSQL_EXPIRE = 57609
const MYSQL_ACCOUNT = 57610
const MYSQL_UNLOCK = 57611
const MYSQL_DAY = 57612
const MYSQL_NEVER = 57613
const MYSQL_SECOND = 57614
const MYSQL_ASCII = 57615
const MYSQL_COALESCE = 57616
const MYSQL_COLLATION = 57617
const MYSQL_HOUR = 57618
const MYSQL_MICROSECOND = 57619
const MYSQL_MINUTE = 57620
const MYSQL_MONTH = 57621
const MYSQL_QUARTER = 57622
const MYSQL_REPEAT = 57623
const MYSQL_REVERSE = 57624
const MYSQL_ROW_COUNT = 57625
const MYSQL_WEEK = 57626
const MYSQL_REVOKE = 57627
const MYSQL_FUNCTION = 57628
const MYSQL_PRIVILEGES = 57629
const MYSQL_TABLESPACE = 57630
const MYSQL_EXECUTE = 57631
const MYSQL_SUPER = 57632
const MYSQL_GRANT = 57633
const MYSQL_OPTION = 57634
const MYSQL_REFERENCES = 57635
const MYSQL_REPLICATION = 57636
const MYSQL_SLAVE = 57637
const MYSQL_CLIENT = 57638
const MYSQL_USAGE = 57639
const MYSQL_RELOAD = 57640
const MYSQL_FILE = 57641
const MYSQL_TEMPORARY = 57642
const MYSQL_ROUTINE = 57643
const MYSQL_EVENT = 57644
const MYSQL_SHUTDOWN = 57645
const MYSQL_NULLX = 57646
const MYSQL_AUTO_INCREMENT = 57647
const MYSQL_APPROXNUM = 57648
const MYSQL_SIGNED = 57649
const MYSQL_UNSIGNED = 57650
const MYSQL_ZEROFILL = 57651
const MYSQL_USER = 57652
const MYSQL_IDENTIFIED = 57653
const MYSQL_CIPHER = 57654
const MYSQL_ISSUER = 57655
const MYSQL_X509 = 57656
const MYSQL_SUBJECT = 57657
const MYSQL_SAN = 57658
const MYSQL_REQUIRE = 57659
const MYSQL_SSL = 57660
const MYSQL_NONE = 57661
const MYSQL_PASSWORD = 57662
const MYSQL_MAX_QUERIES_PER_HOUR = 57663
const MYSQL_MAX_UPDATES_PER_HOUR = 57664
const MYSQL_MAX_CONNECTIONS_PER_HOUR = 57665
const MYSQL_MAX_USER_CONNECTIONS = 57666
const MYSQL_FORMAT = 57667
const MYSQL_VERBOSE = 57668
const MYSQL_CONNECTION = 57669
const MYSQL_LOAD = 57670
const MYSQL_INFILE = 57671
const MYSQL_TERMINATED = 57672
const MYSQL_OPTIONALLY = 57673
const MYSQL_ENCLOSED = 57674
const MYSQL_ESCAPED = 57675
const MYSQL_STARTING = 57676
const MYSQL_LINES = 57677
const MYSQL_DATABASES = 57678
const MYSQL_TABLES = 57679
const MYSQL_EXTENDED = 57680
const MYSQL_FULL = 57681
const MYSQL_PROCESSLIST = 57682
const MYSQL_FIELDS = 57683
const MYSQL_COLUMNS = 57684
const MYSQL_OPEN = 57685
const MYSQL_ERRORS = 57686
const MYSQL_WARNINGS = 57687
const MYSQL_INDEXES = 57688
const MYSQL_SCHEMAS = 57689
const MYSQL_NAMES = 57690
const MYSQL_GLOBAL = 57691
const MYSQL_SESSION = 57692
const MYSQL_ISOLATION = 57693
const MYSQL_LEVEL = 57694
const MYSQL_READ = 57695
const MYSQL_WRITE = 57696
const MYSQL_ONLY = 57697
const MYSQL_REPEATABLE = 57698
const MYSQL_COMMITTED = 57699
const MYSQL_UNCOMMITTED = 57700
const MYSQL_SERIALIZABLE = 57701
const MYSQL_LOCAL = 57702
const MYSQL_EXCEPT = 57703
const MYSQL_CURRENT_TIMESTAMP = 57704
const MYSQL_DATABASE = 57705
const MYSQL_CURRENT_TIME = 57706
const MYSQL_LOCALTIME = 57707
const MYSQL_LOCALTIMESTAMP = 57708
const MYSQL_UTC_DATE = 57709
const MYSQL_UTC_TIME = 57710
const MYSQL_UTC_TIMESTAMP = 57711
const MYSQL_REPLACE = 57712
const MYSQL_CONVERT = 57713
const MYSQL_SEPARATOR = 57714
const MYSQL_CURRENT_DATE = 57715
const MYSQL_CURRENT_USER = 57716
const MYSQL_CURRENT_ROLE = 57717
const MYSQL_SECOND_MICROSECOND = 57718
const MYSQL_MINUTE_MICROSECOND = 57719
const MYSQL_MINUTE_SECOND = 57720
const MYSQL_HOUR_MICROSECOND = 57721
const MYSQL_HOUR_SECOND = 57722
const MYSQL_HOUR_MINUTE = 57723
const MYSQL_DAY_MICROSECOND = 57724
const MYSQL_DAY_SECOND = 57725
const MYSQL_DAY_MINUTE = 57726
const MYSQL_DAY_HOUR = 57727
const MYSQL_YEAR_MONTH = 57728
const MYSQL_SQL_TSI_HOUR = 57729
const MYSQL_SQL_TSI_DAY = 57730
const MYSQL_SQL_TSI_WEEK = 57731
const MYSQL_SQL_TSI_MONTH = 57732
const MYSQL_SQL_TSI_QUARTER = 57733
const MYSQL_SQL_TSI_YEAR = 57734
const MYSQL_SQL_TSI_SECOND = 57735
const MYSQL_SQL_TSI_MINUTE = 57736
const MYSQL_RECURSIVE = 57737
const MYSQL_CONFIG = 57738
const MYSQL_MATCH = 57739
const MYSQL_AGAINST = 57740
const MYSQL_BOOLEAN = 57741
const MYSQL_LANGUAGE = 57742
const MYSQL_WITH = 57743
const MYSQL_QUERY = 57744
const MYSQL_EXPANSION = 57745
const MYSQL_ADDDATE = 57746
const MYSQL_BIT_AND = 57747
const MYSQL_BIT_OR = 57748
const MYSQL_BIT_XOR = 57749
const MYSQL_CAST = 57750
const MYSQL_COUNT = 57751
const MYSQL_APPROX_COUNT_DISTINCT = 57752
const MYSQL_APPROX_PERCENTILE = 57753
const MYSQL_CURDATE = 57754
const MYSQL_CURTIME = 57755
const MYSQL_DATE_ADD = 57756
const MYSQL_DATE_SUB = 57757
const MYSQL_EXTRACT = 57758
const MYSQL_GROUP_CONCAT = 57759
const MYSQL_MAX = 57760
const MYSQL_MID = 57761
const MYSQL_MIN = 57762
const MYSQL_NOW = 57763
const MYSQL_POSITION = 57764
const MYSQL_SESSION_USER = 57765
const MYSQL_STD = 57766
const MYSQL_STDDEV = 57767
const MYSQL_STDDEV_POP = 57768
const MYSQL_STDDEV_SAMP = 57769
const MYSQL_SUBDATE = 57770
const MYSQL_SUBSTR = 57771
const MYSQL_SUBSTRING = 57772
const MYSQL_SUM = 57773
const MYSQL_SYSDATE = 57774
const MYSQL_SYSTEM_USER = 57775
const MYSQL_TRANSLATE = 57776
const MYSQL_TRIM = 57777
const MYSQL_VARIANCE = 57778
const MYSQL_VAR_POP = 57779
const MYSQL_VAR_SAMP = 57780
const MYSQL_AVG = 57781
const MYSQL_ROW = 57782
const MYSQL_OUTFILE = 57783
const MYSQL_HEADER = 57784
const MYSQL_MAX_FILE_SIZE = 57785
const MYSQL_FORCE_QUOTE = 57786
const MYSQL_UNUSED = 57787
