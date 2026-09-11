# Step-by-Step AWS Setup Guide

This guide will walk you through setting up AWS from complete scratch, configuring your local machine, and preparing to deploy your CloudBoard infrastructure.

> [!WARNING]
> **Cost Warning:** While creating an AWS account is free and many resources fall under the Free Tier, deploying resources like EC2, RDS, and Load Balancers will eventually cost money. Always follow the cleanup steps when you are done learning for the day.

---

## Phase 1: Creating Your AWS Account

If you don't have an AWS account yet, start here.

1. **Go to AWS**: Navigate to [aws.amazon.com](https://aws.amazon.com/) and click **Create an AWS Account**.
2. **Provide Email**: Enter your email address and a name for your account. Verify your email with the code sent to you.
3. **Password**: Create a strong password.
4. **Contact Information**: Fill out your contact details (Personal account is fine).
5. **Billing**: You MUST provide a credit or debit card. AWS requires this to verify your identity and to charge you if you exceed the Free Tier.
6. **Identity Verification**: Verify your phone number via text or voice call.
7. **Support Plan**: Select the **Basic Support - Free** plan.
8. **Login**: Go to the AWS Management Console and sign in as the **Root User** using the email and password you just created.

---

## Phase 2: Securing Your Account (Don't use the Root User!)

Using the Root User for daily tasks is a huge security risk. If someone steals your Root User credentials, they have total control over your account. We need to create an "IAM User" for you to use instead.

1. In the AWS Management Console search bar at the top, type **IAM** and click on it.
2. On the left sidebar, click **Users**, then click the **Create user** button.
3. **User details**: 
   - User name: Type `cloudboard-admin` (or your name).
   - Check the box: **Provide user access to the AWS Management Console**.
   - Select **I want to create an IAM user**.
   - Console password: Auto-generated or Custom. Uncheck "Users must create a new password at next sign-in" if you want to keep the custom one.
   - Click **Next**.
4. **Permissions**:
   - Select **Attach policies directly**.
   - In the search bar, type `AdministratorAccess`.
   - Check the box next to **AdministratorAccess** (This gives the user full permissions, which is necessary for Terraform).
   - Click **Next**, then **Create user**.
5. **Save Credentials**: Copy the Console sign-in URL, Username, and Password. 
6. **Log Out** of your Root account, and **Log In** using the new `cloudboard-admin` user and the URL you just copied.

---

## Phase 3: Getting Your Access Keys

To allow Terraform and the AWS CLI on your Mac to talk to AWS, you need Access Keys.

1. Logged in as `cloudboard-admin`, go back to the **IAM** dashboard.
2. Click **Users** on the left, then click on your `cloudboard-admin` user.
3. Click on the **Security credentials** tab.
4. Scroll down to the **Access keys** section and click **Create access key**.
5. Select **Command Line Interface (CLI)**, check the confirmation box, and click **Next**.
6. Give it a description tag (e.g., `MacBook CLI`) and click **Create access key**.
7. **IMPORTANT**: You will see an **Access key ID** and a **Secret access key**. 
   - Click **Download .csv file** and save it somewhere extremely safe.
   - You will *never* be able to see the Secret Access Key again once you close this page. If you lose it, you have to delete it and create a new one.

---

## Phase 4: Installing and Configuring the AWS CLI

Now we move to your Mac terminal.

1. **Install the AWS CLI**:
   Open your Mac terminal and run these commands to download and install the official AWS CLI:
   ```bash
   curl "https://awscli.amazonaws.com/AWSCLIV2.pkg" -o "AWSCLIV2.pkg"
   sudo installer -pkg AWSCLIV2.pkg -target /
   rm AWSCLIV2.pkg
   ```
2. **Verify Installation**:
   ```bash
   aws --version
   ```
3. **Configure the CLI**:
   Run the following command:
   ```bash
   aws configure
   ```
   It will prompt you for 4 pieces of information. Use the keys you generated in Phase 3:
   - **AWS Access Key ID**: Paste your Access Key ID.
   - **AWS Secret Access Key**: Paste your Secret Access Key.
   - **Default region name**: Type `us-east-1` (or whichever region is closest to you, like `eu-west-1`).
   - **Default output format**: Type `json`.

4. **Verify Connection**:
   Run this command to prove your Mac is talking to your AWS account:
   ```bash
   aws sts get-caller-identity
   ```
   You should see a JSON response with your Account ID and the `cloudboard-admin` ARN.

---

## Phase 5: Installing Terraform

Terraform is the tool that reads our `main.tf` files and automatically builds the AWS infrastructure.

1. **Install Terraform via Homebrew**:
   ```bash
   brew tap hashicorp/tap
   brew install hashicorp/tap/terraform
   ```
2. **Verify Installation**:
   ```bash
   terraform -v
   ```

---

## Phase 6: Initializing Terraform

Now you are ready to use Terraform with the CloudBoard project!

1. Navigate to the Terraform development directory in your terminal:
   ```bash
   cd infra/terraform/env/dev
   ```
2. Initialize Terraform (this downloads the AWS provider plugins):
   ```bash
   terraform init
   ```
3. See what Terraform *plans* to do:
   ```bash
   terraform plan
   ```
   *(Note: Right now, our `main.tf` is just a skeleton. It won't build much until we fill out the modules, but this proves it can connect to AWS!)*
4. If you want to actually build the resources, you would run:
   ```bash
   terraform apply
   ```
   (Type `yes` when prompted).
5. **ALWAYS CLEAN UP**: When you are done learning for the day, destroy the resources so you don't get billed:
   ```bash
   terraform destroy
   ```
   (Type `yes` when prompted).
