
import java.util.*;

public class EmailAddresses {
    private LinkedList<Messages> Sent = new LinkedList<Messages>();
    private LinkedList<Messages> Received = new LinkedList<Messages>();
    private String OwnerAddress;

    public String getOwnerAddress() {
        return this.OwnerAddress;
    }

    public void setReceived(Messages msg) {
        this.Received.add(msg);
    }

    public EmailAddresses(String Address) {
        this.OwnerAddress = Address;
    }

}
