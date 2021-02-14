package SE421.classes;

import java.io.*;
import java.util.Scanner;

/**
 * assignment
 */
public class assignment {
    public static void main(String[] args) throws IOException {
        File txt = new File("result.txt");
        FileWriter fw = new FileWriter(txt);
        PrintWriter pw = new PrintWriter(fw);
        double r1, r2, r3;
        double x, y, a, b, c, d;
        Scanner in = new Scanner(System.in);
        // addition
        System.out.println("please enter two decimal numbers");
        System.out.printf("Number 1:");
        x = in.nextDouble();
        System.out.printf("Number 2:");
        y = in.nextDouble();
        r1 = x + y;
        System.out.println("Addition result is " + r1);
        // multiplication
        System.out.println("please enter two decimal numbers");
        System.out.printf("Number 1:");
        a = in.nextDouble();
        System.out.printf("Number 2:");
        b = in.nextDouble();
        r2 = a * b;
        System.out.println("Multiplication result is " + r2);
        // division
        System.out.println("please enter two decimal numbers");
        System.out.printf("Number 1:");
        c = in.nextDouble();
        do {
            System.out.printf("Number 2:");
            d = in.nextDouble();
            if (d == 0)
                System.out.println("Invalid input please retry");

        } while (d == 0);
        r3 = c / d;
        System.out.println("Division result is " + r3);
        pw.println("The addition result of " + x + " and " + y + " is " + r1);
        pw.println("The multiplication result of " + a + " and " + b + " is " + r2);
        pw.println("The division result of " + c + " and " + d + " is " + r3);
        pw.close();

    }

}