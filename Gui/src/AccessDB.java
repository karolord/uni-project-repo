import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.SQLException;
import java.sql.Statement;

public class AccessDB {
    public static void main(String[] args) {
        try {
            String x = "1A";
            Connection con = DriverManager.getConnection("jdbc:mysql://localhost:3306/airlinereservationdb", "root",
                    "root");
            System.out.println("connected");
            Statement st = con.createStatement();
            st.executeUpdate("UPDATE airlinereservationdb.seat SET s_Reserved = 0 WHERE s_number ='" + x + "';");
        } catch (SQLException ex) {
            System.out.println(ex.getMessage());
        }
    }
}
