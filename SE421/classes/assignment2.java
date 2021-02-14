package SE421.classes;

import java.io.*;
import java.util.Scanner;

public class assignment2 {
    public static void main(String[] args) throws IOException {
        Result r1 = new Result();
        Inputs i1 = new Inputs();
        Inputs i2 = new Inputs();
        System.out.println("please enter two decimal numbers");
        i1.value = Rvalue();
        i2.value = Rvalue();
        r1.value = Addition(i1.value, i2.value);
        Print(r1.value);
        i1.value = Rvalue();
        i2.value = Rvalue();
        r1.value = Multiplication(i1.value, i2.value);
        Print(r1.value);
        i1.value = Rvalue();
        i2.value = Rvalue();
        r1.value = Division(i1.value, i2.value);
        Print(r1.value);
        text(r1.value);

    }

    public static class Result {
        double value;

    }

    public static class Inputs {
        double value;
    }

    public static void Print(double a) {
        System.out.println("The Result is " + a);

    }

    public static double Addition(double a, double b) {
        return a + b;
    }

    public static double Multiplication(double a, double b) {
        return a * b;
    }

    public static double Division(double a, double b) {
        Scanner in = new Scanner(System.in);

        while (b == 0) {
            System.out.println("error invalid argument retry");
            System.out.printf("Number :");
            b = in.nextDouble();
        }
        return a / b;
    }

    // input value
    public static double Rvalue() {
        Scanner in = new Scanner(System.in);
        System.out.printf("Number :");
        double a = in.nextDouble();
        return a;

    }

    public static void text(double a) throws IOException {
        File txt = new File("result.txt");
        FileWriter fw = new FileWriter(txt);
        PrintWriter pw = new PrintWriter(fw);
        pw.println("The division result is " + a);
        pw.close();

    }

}
