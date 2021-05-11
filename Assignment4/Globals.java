import java.util.*;
public class Globals {
  
  //public class main
  public static LinkedList<MessageInterface> allMessages = new LinkedList<MessageInterface>();
  public static LinkedList<AddressInterface> allEmails = new LinkedList<AddressInterface>();
  public static EmailAddresses CurrentEmail;
  
  //public class Ui
  public static Scanner scanner = new Scanner(System.in);
  
  //public class Messages
  public static final String Title;
  public static final String Description;
  public static final LinkedList<AddressInterface> SentTo = new LinkedList<AddressInterface>();
  public static final LinkedList<AddressInterface> CC = new LinkedList<AddressInterface>();
  
  //public class MessageInterface
  public static LinkedList<AddressInterface> getSentTo();
  public static LinkedList<AddressInterface> getCC();
  
  //public class EmailAddresses
  public static String OwnerAddress;
  public static LinkedList<MessageInterface> Sent = new LinkedList<MessageInterface>();
  public static LinkedList<MessageInterface> Received = new LinkedList<MessageInterface>();
  
}
