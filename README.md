
# stop-start-rds-golang

Function to start or stop RDS in golang.

## Roles

The lambda function needs to have the **AmazonRDSFullAccess** role

## Cloudwatch event:

    {  
	   "type": "START",
       "dbName": "dbName" 
    }


### Parameters
**type**: START/STOP
**dbName**: the RDS identifier
