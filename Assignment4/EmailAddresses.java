
import java.util.*;

public class EmailAddresses {
    private String OwnerAddress;
    private LinkedList<Messages> Sent = new LinkedList<Messages>();
    private LinkedList<Messages> Received = new LinkedList<Messages>();

    public LinkedList<Messages> getReceived() {
        return this.Received;
    }

    public LinkedList<Messages> getSent() {
        return this.Sent;
    }

    public void setSent(Messages Sent) {
        this.Sent.add(Sent);
    }

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
