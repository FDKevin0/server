import os

MYSQL_HOST = os.environ.get("MYSQL_HOST", "127.0.0.1")
MYSQL_PORT = int(os.environ.get("MYSQL_PORT", 3306))
MYSQL_USER = os.environ.get("MYSQL_USER", "user")
MYSQL_PASS = os.environ.get("MYSQL_PASS", "password")
MYSQL_DB = os.environ.get("MYSQL_DB", "bangumi")
