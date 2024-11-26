function base32Decode(input) {
  const alphabet = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ234567';
  let buffer = 0, bitsLeft = 0, output = [];

  for (let char of input) {
    const val = alphabet.indexOf(char.toUpperCase());
    if (val === -1) continue; // Ignore invalid characters
    buffer = (buffer << 5) | val;
    bitsLeft += 5;

    if (bitsLeft >= 8) {
      bitsLeft -= 8;
      output.push((buffer >> bitsLeft) & 0xFF);
      buffer &= (1 << bitsLeft) - 1;
    }
  }
  return Uint8Array.from(output);
}

function AWSAccountFromAWSKeyID(AWSKeyID) {
  if (AWSKeyID.length <= 4) return "Invalid Key ID";
  let trimmed = AWSKeyID.slice(4);
  let decoded = base32Decode(trimmed);
  let bigIntValue = BigInt('0x' + Array.from(decoded.slice(0, 6)).map(b => b.toString(16).padStart(2, '0')).join(''));
  let mask = BigInt('0x7fffffffff80');
  return ((bigIntValue & mask) >> 7n).toString().padStart(12, '0');
}

// Example Usage
let awsKeyID = "ABCDABCDEFGHIJKLMNOPQRSTUVWX";
console.log("AWS Account ID:", AWSAccountFromAWSKeyID(awsKeyID));

