using System;
using System.Collections.Generic;

class AwsAccountIdExtractor
{
    static void Main(string[] args)
    {
        // Example AWS Access Key ID
        string awsAccessKeyId = "AKIAEXAMPLE123456";

        // Extract AWS Account ID
        string accountId = GetAwsAccountId(awsAccessKeyId);
        if (accountId == null)
        {
            Console.WriteLine("Invalid AWS Access Key ID.");
        }
        else
        {
            Console.WriteLine($"AWS Account ID: {accountId}");
        }

        // Identify Resource Type
        string resourceType = GetResourceType(awsAccessKeyId.Substring(0, 4));
        Console.WriteLine($"Resource Type: {resourceType}");
    }

    static string GetAwsAccountId(string awsKeyId)
    {
        if (awsKeyId.Length <= 4)
            return null;

        string trimmedKey = awsKeyId.Substring(4);
        byte[] decoded = Base32Decode(trimmedKey);

        if (decoded.Length < 6)
            return null;

        // Extract the first 6 bytes
        byte[] extractedBytes = new byte[6];
        Array.Copy(decoded, extractedBytes, 6);

        // Convert to integer and shift right 7 bits
        long accountId = (long)(
            (ulong)extractedBytes[0] << 40 |
            (ulong)extractedBytes[1] << 32 |
            (ulong)extractedBytes[2] << 24 |
            (ulong)extractedBytes[3] << 16 |
            (ulong)extractedBytes[4] << 8 |
            (ulong)extractedBytes[5]
        ) >> 7;

        // Return as 12-digit zero-padded string
        return accountId.ToString("D12");
    }

    static string GetResourceType(string prefix)
    {
        return prefix switch
        {
            "AKIA" => "IAM User",
            "ASIA" => "Temporary Security Credential",
            "AIDA" => "Federated User",
            "ANPA" => "S3 Service Role",
            _ => "Unknown Resource Type",
        };
    }

    static byte[] Base32Decode(string input)
    {
        const string alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567";
        int buffer = 0, bitsLeft = 0;
        var output = new List<byte>();

        foreach (char c in input.ToUpperInvariant())
        {
            int value = alphabet.IndexOf(c);
            if (value == -1) continue; // Ignore invalid characters

            buffer = (buffer << 5) | value;
            bitsLeft += 5;

            if (bitsLeft >= 8)
            {
                bitsLeft -= 8;
                output.Add((byte)(buffer >> bitsLeft));
                buffer &= (1 << bitsLeft) - 1;
            }
        }

        return output.ToArray();
    }
}

