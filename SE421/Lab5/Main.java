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
        add ye = new add();
        yes.staroperation(n, ye);
        ex.save(yes);

    }
}

interface Saver {
    double getresult();
}

interface Operator {
    double operation(double input1, double input2);

}

class Stararr implements Saver {
    double value;
    double result;
    Scanner scanner = new Scanner(System.in);

    public void staroperation(int n, Operator o) {
        for (int i = 0; i < n; i++) {
            System.out.println("Value:");
            value = scanner.nextDouble();
            result = o.operation(result, value);
        }
        System.out.println(result);
    }
    public double getresult(){
        return result;
    }
}

class add implements Operator, Saver {
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

class Subtract implements Operator, Saver {
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

class multiply implements Operator, Saver {
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

class divide implements Operator, Saver {
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

    public void save(Saver s) throws IOException {
        File txt = new File("result.txt");
        FileWriter fw = new FileWriter(txt);
        PrintWriter pw = new PrintWriter(fw);
        pw.println("The result is " + s.getresult());
        pw.close();
    }

}