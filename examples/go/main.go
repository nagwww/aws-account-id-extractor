package main

import (
	"fmt"
	"strings"
)

// base32Decode decodes a Base32 encoded string.
func base32Decode(input string) ([]byte, error) {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"
	var buffer uint32
	var bitsLeft uint8
	output := make([]byte, 0, len(input)*5/8)

	for _, char := range input {
		index := strings.IndexRune(alphabet, char)
		if index == -1 {
			continue // Ignore invalid characters
		}

		buffer = (buffer << 5) | uint32(index)
		bitsLeft += 5

		if bitsLeft >= 8 {
			bitsLeft -= 8
			output = append(output, byte(buffer>>bitsLeft))
			buffer &= (1 << bitsLeft) - 1
		}
	}

	return output, nil
}

// getAwsAccountId extracts the AWS Account ID from an AWS Access Key ID.
func getAwsAccountId(awsKeyId string) (string, error) {
	if len(awsKeyId) <= 4 {
		return "", fmt.Errorf("invalid Key ID")
	}

	trimmedKey := awsKeyId[4:]
	decoded, err := base32Decode(trimmedKey)
	if err != nil {
		return "", fmt.Errorf("failed to decode base32: %v", err)
	}

	if len(decoded) < 6 {
		return "", fmt.Errorf("decoded data too short")
	}

	// Extract the first 6 bytes
	extractedBytes := decoded[:6]

	// Convert bytes to an integer, then shift right 7 bits
	accountId := (uint64(extractedBytes[0]) << 40) |
		(uint64(extractedBytes[1]) << 32) |
		(uint64(extractedBytes[2]) << 24) |
		(uint64(extractedBytes[3]) << 16) |
		(uint64(extractedBytes[4]) << 8) |
		uint64(extractedBytes[5])
	accountId >>= 7

	// Return the 12-digit zero-padded account ID
	return fmt.Sprintf("%012d", accountId), nil
}

// getResourceType identifies the resource type from the AWS Access Key ID prefix.
func getResourceType(prefix string) string {
	prefixMapping := map[string]string{
		"AKIA": "IAM User",
		"ASIA": "Temporary Security Credential",
		"AIDA": "Federated User",
		"ANPA": "S3 Service Role",
	}

	if resourceType, exists := prefixMapping[prefix]; exists {
		return resourceType
	}
	return "Unknown Resource Type"
}

func main() {
	// Example AWS Access Key ID
	awsAccessKeyId := "AKIAEXAMPLE123456"

	// Extract AWS Account ID
	accountId, err := getAwsAccountId(awsAccessKeyId)
	if err != nil {
		fmt.Printf("Error extracting AWS Account ID: %v\n", err)
		return
	}
	fmt.Printf("AWS Account ID: %s\n", accountId)

	// Identify Resource Type
	resourceType := getResourceType(awsAccessKeyId[:4])
	fmt.Printf("Resource Type: %s\n", resourceType)
}

