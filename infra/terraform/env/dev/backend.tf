# Wait to uncomment this until AFTER you have run the bootstrap and created your bucket!
# terraform {
#   backend "s3" {
#     bucket         = "cloudboard-tf-state-2dc080b0"
#     key            = "dev/terraform.tfstate"
#     region         = "us-east-1"
#     dynamodb_table = "cloudboard-tf-locks"
#     encrypt        = true
#   }
# }
