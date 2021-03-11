package SE421.Lab5;

import java.io.*;
import java.util.Scanner;

public class Main {

    public static void main(String[] args) throws IOException {
        Scanner scanner = new Scanner(System.in);
        Exporter ex = new Exporter();
        System.out.println("please enter the numbers of inputs");
        int n = scanner.nextInt();
        Stararr yes = new Stararr();
        Subtract ye = new Subtract();
        yes.staroperation(n, ye);
        ex.save(yes);

    }
}

interface hasresult  {
    double getresult();
}

interface premiativeoperator{
    double operation(double input1, double input2);

}

class Stararr implements hasresult {
    double value;
    double result;
    Scanner scanner = new Scanner(System.in);

    public void staroperation(int n, premiativeoperator o) {
        System.out.printf("value:");
        result = scanner.nextDouble();
        for (int i = 0; i < n - 1; i++) {
            System.out.printf("Value:");
            value = scanner.nextDouble();
            result = o.operation(result, value);
        }
        System.out.println(result);
    }

    public double getresult() {
        return result;
    }
}

class Add implements premiativeoperator,hasresult {
    double value1;
    double value2;
    double result;

    public double operation(double input1, double input2) {
        value1 = input1;
        value2 = input2;
        result = input1 + input2;
        return result;
    }

    public double getresult() {
        return result;
    }
}

class Subtract implements premiativeoperator,hasresult {
    double value1;
    double value2;
    double result;

    public double operation(double input1, double input2) {
        value1 = input1;
        value2 = input2;
        result = input1 - input2;
        return result;
    }

    public double getresult() {
        return result;
    }
}

class Multiply implements premiativeoperator,hasresult {
    double value1;
    double value2;
    double result;

    public double operation(double input1, double input2) {
        value1 = input1;
        value2 = input2;
        result = input1 * input2;
        return result;
    }

    public double getresult() {
        return result;
    }
}

class Divide implements premiativeoperator,hasresult {
    double value1;
    double value2;
    double result;

    public double operation(double input1, double input2) {
        value1 = input1;
        value2 = input2;
        result = input1 / input2;
        return result;
    }

    public double getresult() {
        return result;
    }
}

class Exporter {

    public void save(hasresult s) throws IOException {
        File txt = new File("result.txt");
        FileWriter fw = new FileWriter(txt);
        PrintWriter pw = new PrintWriter(fw);
        pw.println("The result is " + s.getresult());
        pw.close();
    }

}