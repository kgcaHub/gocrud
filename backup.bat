del crudDB
SQLCMD -E -S . -Q "BACKUP DATABASE crudDB TO DISK ='%cd%\crudDB'"