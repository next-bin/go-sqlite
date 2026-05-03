# example usage
#
# $ ./bisect.sh ccgo s390
# modernc.org/ccgo/v4/lib|s390x|2025-02-19 20:01:01+01:00|7f67dcd79feb54fb6c270ebd986120c9e7bc41ca|linux|s390x|PASS|go1.23.5
# modernc.org/ccgo/v4/lib|s390x|2025-03-05 02:19:12+01:00|377cffa218bcac8e8abe2e18d9df665fa6e930e9|linux|s390x|FAIL|go1.23.5
# $

DB=.exclude/results2.db
if [ ! -f $DB ] ; then make db ; fi
sqlite3 $DB "select * from results where import_path like '%$1%' and builder = '$2' and pass='PASS' order by date desc limit 1;"
D=$(sqlite3 $DB "select date from results where import_path like '%$1%' and builder = '$2' and pass='PASS' order by date desc limit 1;")
sqlite3 $DB "select * from results where import_path like '%$1%' and builder = '$2' and pass='FAIL' and date > '$D' order by date limit 1;"
