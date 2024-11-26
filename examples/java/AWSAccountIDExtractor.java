import java.math.BigInteger;
import java.util.Base64;

public class AWSAccountIDExtractor {

    private static final String ALPHABET = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567";

    public static byte[] base32Decode(String input) {
        int buffer = 0, bitsLeft = 0;
        byte[] output = new byte[(int) Math.ceil(input.length() * 5 / 8)];
        int index = 0;

        for (char c : input.toCharArray()) {
            int value = ALPHABET.indexOf(Character.toUpperCase(c));
            if (value == -1) continue; // Ignore invalid characters
            buffer = (buffer << 5) | value;
            bitsLeft += 5;

            if (bitsLeft >= 8) {
                bitsLeft -= 8;
                output[index++] = (byte) (buffer >> bitsLeft);
                buffer &= (1 << bitsLeft) - 1;
            }
        }
        return java.util.Arrays.copyOf(output, index);
    }

    public static String getAWSAccountID(String awsKeyID) {
        if (awsKeyID.length() <= 4) return "Invalid Key ID";
        String trimmed = awsKeyID.substring(4);
        byte[] decoded = base32Decode(trimmed);
        byte[] accountBytes = java.util.Arrays.copyOfRange(decoded, 0, 6);
        BigInteger accountID = new BigInteger(1, accountBytes).shiftRight(7);
        return String.format("%012d", accountID);
    }

    public static void main(String[] args) {
        String awsKey = "ABCDABCDEFGHIJKLMNOPQRSTUVWX";
        System.out.println("AWS Account ID: " + getAWSAccountID(awsKey));
    }
}
