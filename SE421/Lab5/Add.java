package SE421.Lab5;

public class Add {
    double input1;
    double input2;
    double result;

    public double add(double value1, double value2) {
        input1 = value1;
        input2 = value2;
        result = value1 + value2;
        return result;
    }

    public double getresult() {
        return result;
    }

}
