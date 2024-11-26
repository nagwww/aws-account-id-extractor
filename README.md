# **AWS Account ID Extractor**

This repository includes examples in Python, JavaScript, Java, C#, and Go for programmatic validation and processing of AWS credentials. Implementation examples to extract AWS Account IDs and identify resource types from AWS Access Key IDs

---

## **Overview**

AWS Access Key IDs often include encoded information about the AWS account and resource type. This project demonstrates how to:

1. Extract the **AWS Account ID** from an AWS Access Key ID.  
2. Identify the **resource type** based on the prefix of the AWS Access Key ID.

The examples are available in the following languages:

* **Python**  
* **JavaScript**  
* **Java**  
* **C\#**  
* **Go**

---

## **Features**

* Extract **AWS Account ID** from AWS Access Key IDs.  
* Identify the **resource type** based on AWS Access Key ID prefixes (e.g., IAM User, Service Role).

---

## **Language Implementations**

### **1\. Python**

The Python implementation is available as a PyPI package.

#### **Setup**

Install the library:
 
```
pip install aws_access_key_id
```

#### **Usage**

```
from aws_access_key_id import get_aws_account_id, get_resource_type

aws_access_key_id = "AKIAEXAMPLE123456"

# Extract AWS Account ID  
account_id = get_aws_account_id(aws_access_key_id)  
print(f"AWS Account ID: {account_id}")

# Identify resource type  
resource_type = get_resource_type(aws_access_key_id[:4])
print(f"Resource Type: {resource_type}")
```

### **2\. JavaScript**

The JavaScript implementation uses plain Node.js.

#### **Setup**

Install dependencies:


```
npm install
```

#### **Usage**
 
```
node aws_account_id_extractor.js
```

### **3\. Java**

The Java implementation is structured as a Maven project.

#### **Setup**

Compile the project:
  
```
javac AWSAccountIDExtractor.java
```

#### **Usage**

```
java AWSAccountIDExtractor
```

### **4\. C\#**

The C\# implementation is part of a .NET Core project.

#### **Setup**

Build the project:

```
dotnet build
```

#### **Usage**
  
```
csc AwsAccountIdExtractor.cs
```

```
./AwsAccountIdExtractor

AWS Account ID: 123456789012
Resource Type: IAM User
```

### **5\. Go**

The Go implementation provides a simple utility function.

#### **Setup**

Install dependencies:
 
```
brew install go
```

#### **Usage**
 
```
cd examples/go
go run main.go
```

---

## **AWS Access Key ID Structure**

AWS Access Key IDs typically begin with a 4-character prefix indicating the resource type:

| Prefix | Resource Type |
| ----- | ----- |
| AKIA | IAM User |
| ASIA | Temporary Security Credential |
| AIDA | Federated User |
| ANPA | S3 Service Role |

---

## **Contributing**

Contributions are welcome\! If you'd like to add support for another programming language or enhance existing implementations:

1. Fork the repository.  
2. Create a feature branch.  
3. Submit a pull request.

---

## **License**

This project is licensed under the MIT License.

---

## **Acknowledgments**


* Thanks AWS for their comprehensive documentation, which helped guide this implementation.
* Special Thanks Tal Be'ery ( [@ZenGo (KZen](https://x.com/TalBeerySec) ) https://medium.com/@TalBeerySec/a-short-note-on-aws-key-id-f88cc4317489


