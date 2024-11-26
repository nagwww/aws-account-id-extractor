from aws_access_key_id import get_aws_account_id, get_resource_type

# Example AWS Access Key ID
aws_access_key_id = "AKIAEXAMPLE123456"

# Extract account ID
account_id = get_aws_account_id(aws_access_key_id)
print(f"AWS Account ID: {account_id}")

# Identify resource type
resource_type = get_resource_type(aws_access_key_id[:4])
print(f"Resource Type: {resource_type}")

