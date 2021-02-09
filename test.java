import java.security.SecureRandom;

public class test {
    public static void main(String[] args) {
        int[] array = new int[5];
        SecureRandom test = new SecureRandom();
        for (int i = 0; i < array.length; i++) {
            array[i] = test.nextInt(5);
        }
        for (int i : array) {
            System.out.println(array[i]);
        }
    }

}