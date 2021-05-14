import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.ResultSet;
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
            ResultSet rs = st.executeQuery(
                    "Select s_reserved from airlinereservationdb.seat WHERE s_number = '1C' AND f_code = 'AUI22';");
            rs.next();
            System.out.println(rs.getObject(1));
            st.executeUpdate("UPDATE airlinereservationdb.seat SET s_Reserved = 0 WHERE s_number ='" + x + "';");
        } catch (SQLException ex) {
            System.out.println(ex.getMessage());
        }
    }
}
