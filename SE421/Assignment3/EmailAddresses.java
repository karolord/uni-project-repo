import java.util.*;

public class EmailAddresses implements Em {
    private String OwnerAddress;
    private LinkedList<Messages> Sent = new LinkedList<Messages>();
    private LinkedList<Messages> Received = new LinkedList<Messages>();

    public void setReceived(Messages msg) {
        this.Received.add(msg);
    }

    public EmailAddresses(String Address) {
        this.OwnerAddress = Address;
    }

}
