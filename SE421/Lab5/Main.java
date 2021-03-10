package SE421.Lab5;

import java.util.Scanner;

public class Main {
    Scanner scanner = new Scanner(System.in);

    public static void main(String[] args) {
        Stararr yes = new Stararr();
        Subtract ye = new Subtract();
        yes.staradd(5, ye);
    }
}

/**
 * InnerMain
 */
interface Operator {
    double operation(double input1, double input2);

}

class Stararr {
    double value;
    double result;
    Scanner scanner = new Scanner(System.in);

    public void staradd(int n, Operator o) {
        for (int i = 0; i < n; i++) {
            System.out.println("Value:");
            value = scanner.nextDouble();
            result = o.operation(result, value);
        }
        System.out.println(result);
    }

}

class add implements Operator {
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

class Subtract implements Operator {
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

class multiply implements Operator {
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

class divide implements Operator {
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