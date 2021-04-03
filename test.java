import java.security.SecureRandom;

public class test {
    public static void main(String[] args) {

        for (int row = 0; row < 7; row++) {
            for (int i = 0; i < row; i++) {
                System.out.printf(" ");
            }

            for (int i = 0; i < 13 - (2 * row); i++) {
                System.out.printf("x");
            }
            System.out.println();
        }

        for (int row = 6; row > 0; row--) {
            for (int i = 0; i < row; i++) {
                System.out.printf(" ");
            }

            for (int i = 0; i < 13 - (2 * row); i++) {
                System.out.printf("x");
            }
            System.out.println();
        }

    }
}