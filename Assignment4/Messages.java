import java.util.*;

public class Messages implements MessageInterface {
    private final String Title;
    private final String Description;
    private final LinkedList<AddressInterface> SentTo = new LinkedList<AddressInterface>();
    private final LinkedList<AddressInterface> CC = new LinkedList<AddressInterface>();

    public LinkedList<AddressInterface> getSentTo() {
        return this.SentTo;
    }

    public LinkedList<AddressInterface> getCC() {
        return this.CC;
    }

    public String getTitle() {
        return this.Title;
    }

    public String getDescription() {
        return this.Description;
    }

    public Messages(String Title, String Description, LinkedList<AddressInterface> SentTo,
            LinkedList<AddressInterface> CC) {
        this.Title = Title;
        this.Description = Description;
        for (Iterator i = SentTo.iterator(); i.hasNext();) {
            this.SentTo.add((AddressInterface) i.next());
        }
        for (Iterator i = CC.iterator(); i.hasNext();) {
            this.CC.add((AddressInterface) i.next());
        }
    }
}
