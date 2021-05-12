import java.util.*;

public class EmailAddresses implements AddressInterface {
    private String OwnerAddress;
    private LinkedList<MessageInterface> Sent = new LinkedList<MessageInterface>();
    private LinkedList<MessageInterface> Received = new LinkedList<MessageInterface>();

    public LinkedList<MessageInterface> getReceived() {
        return this.Received;
    }

    public LinkedList<MessageInterface> getSent() {
        return this.Sent;
    }

    public void setSent(MessageInterface Sent) {
        this.Sent.add(Sent);
    }

    public String getOwnerAddress() {
        return this.OwnerAddress;
    }

    public void setReceived(MessageInterface msg) {
        this.Received.add(msg);
    }

    public EmailAddresses(String Address) {
        this.OwnerAddress = Address;
    }

    public void Print() {
        System.out.println(this.OwnerAddress);
    }

}
