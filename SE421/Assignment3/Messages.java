import java.util.*;

public class Messages {
    private final String Title;
    private final String Description;
    private final LinkedList<EmailAddresses> SentTo = new LinkedList<EmailAddresses>();
    private final LinkedList<EmailAddresses> CC = new LinkedList<EmailAddresses>();

    public final LinkedList<EmailAddresses> getSentTo() {
        return this.SentTo;
    }

    public final LinkedList<EmailAddresses> getCC() {
        return this.CC;
    }

    public String getTitle() {
        return this.Title;
    }

    public String getDescription() {
        return this.Description;
    }

    public Messages(String Title, String Description, LinkedList<EmailAddresses> SentTo,
            LinkedList<EmailAddresses> CC) {
        this.Title = Title;
        this.Description = Description;
        for (Iterator i = SentTo.iterator(); i.hasNext();) {
            this.SentTo.add((EmailAddresses) i.next());
        }
        for (Iterator i = CC.iterator(); i.hasNext();) {
            this.CC.add((EmailAddresses) i.next());
        }
    }
}
