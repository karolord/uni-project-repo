import java.util.*;

public class Factory {
    public static EmailAddresses GetEmail(String Address, Messages Message) {
        return new EmailAddresses(Address, Message);
    }

    public static Messages GetMessage(String Title, String Description, LinkedList<EmailAddresses> SentTo,
            LinkedList<EmailAddresses> CC) {
        return new Message(Title, Description, SentTo, CC);
    }
}
