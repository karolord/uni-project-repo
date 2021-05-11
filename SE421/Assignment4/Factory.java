import java.util.*;

public class Factory {
    public static EmailAddresses GetEmail(String Address) {
        return new EmailAddresses(Address);
    }

    public static Messages GetMessage(String Title, String Description, LinkedList<EmailAddresses> SentTo,
            LinkedList<EmailAddresses> CC) {
        return new Messages(Title, Description, SentTo, CC);
    }
}
