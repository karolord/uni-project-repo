import java.util.*;

public class Ui {
    public static Scanner scanner = new Scanner(System.in);

    public static void main(String[] args) {
        ComposeMessage();
    }

    public static void startmenu() {
        System.out.println("Welcome to the greatest email client application");
        System.out.println("Enter your email address: ");
    }

    public static void ComposeMessage() {
        System.out.println("Please enter the message Title");
        String title = scanner.nextLine();
        System.out.println("Please enter the message Description");
        String Description = scanner.nextLine();
        LinkedList<EmailAddresses> SentTo = addEmails();
        LinkedList<EmailAddresses> CC = addEmails();
        System.out.println("end");

    }

    public static LinkedList<EmailAddresses> addEmails() {
        LinkedList<EmailAddresses> tmplink = new LinkedList<EmailAddresses>();
        while (true) {
            System.out.println(
                    "Please enter the emails you want to sent the message to one at a time:( or 0 to exit the loop)");
            String email = scanner.next();
            if (email.equals("0"))
                break;
            EmailAddresses emailAddress = Factory.GetEmail(email);
            boolean x = true;
            for (Iterator i = main.allEmails.iterator(); i.hasNext();) {
                EmailAddresses tmp = (EmailAddresses) i.next();
                if (email.equals((tmp.getOwnerAddress()))) {
                    tmplink.add(tmp);
                    x = false;
                }
            }
            if (x) {
                tmplink.add(emailAddress);
                main.allEmails.add(emailAddress);
            }
        }
        return tmplink;
    }
}
