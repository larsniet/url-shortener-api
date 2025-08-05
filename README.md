# How to use the Atlas migration system:

## Apply migrations:

atlas migrate apply --env local --allow-dirty

## Create new migrations:

After updating schema.hcl:
atlas migrate diff migration_name --env local

## Check migration status:

atlas migrate status --env local --allow-dirty
