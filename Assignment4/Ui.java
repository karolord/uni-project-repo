import java.util.*;

public class Ui {

    public static void startmenu() {
        System.out.println("Welcome to the greatest email client application");
        System.out.println("Enter your email address: ");
        String Current = Globals.scanner.nextLine();
        Globals.CurrentEmail = new EmailAddresses(Current);
        testemails();
        PrintInbox();
        Globals.allEmails.add(Globals.CurrentEmail);
        while (true) {
            System.out.println("Please select one of the following");
            System.out.println("1.Compose an email");
            System.out.println("2.Print sent directory");
            System.out.println("3.Print recieved directory");
            System.out.println("4.Exit");
            int selection = Globals.scanner.nextInt();
            Globals.scanner.nextLine();
            switch (selection) {
                case 1:
                    ComposeMessage();
                    break;
                case 2:
                    PrintSentdirectory();
                    break;
                case 3:
                    PrintInbox();
                    break;
                default:
                    return;
            }
        }
    }

    public static void testemails() {
        Globals.allEmails.add(new EmailAddresses("Karo@auis.com"));
        Globals.allEmails.add(new EmailAddresses("Kosar@auis.com"));
        Messages tmp = new Messages("Test email 1", "This email is a test", Globals.allEmails, Globals.allEmails);
        Globals.CurrentEmail.setReceived(tmp);
        Globals.allMessages.add(tmp);
        tmp = new Messages("Test email 2", "This email is a test", Globals.allEmails, Globals.allEmails);
        Globals.CurrentEmail.setReceived(tmp);
        Globals.allMessages.add(tmp);
        tmp = new Messages("Test email 3", "This email is a test", Globals.allEmails, Globals.allEmails);
        Globals.CurrentEmail.setReceived(tmp);
        Globals.allMessages.add(tmp);
    }

    public static void PrintInbox() {
        LinkedList<MessageInterface> inbox = Globals.CurrentEmail.getReceived();
        System.out.println("Here are the messages available in your inbox");
        for (Iterator i = inbox.iterator(); i.hasNext();) {
            Messages tmpmsg = (Messages) i.next();
            System.out.println("title: " + tmpmsg.getTitle());
            System.out.println("Description: " + tmpmsg.getDescription());
            Printemails(tmpmsg);
            System.out.println("-----------------------------");
        }
    }

    public static void PrintSentdirectory() {
        LinkedList<MessageInterface> inbox = Globals.CurrentEmail.getSent();
        System.out.println("Here are the messages available in your sent directory");
        for (Iterator i = inbox.iterator(); i.hasNext();) {
            Messages tmpmsg = (Messages) i.next();
            System.out.println("title: " + tmpmsg.getTitle());
            System.out.println("Description: " + tmpmsg.getDescription());
            Printemails(tmpmsg);
            System.out.println("-----------------------------");
        }
    }

    public static void Printemails(Messages msg) {
        System.out.println("Email is sent to:");
        LinkedList<AddressInterface> tmplink = msg.getSentTo();
        for (Iterator i = tmplink.iterator(); i.hasNext();) {
            EmailAddresses tmpemail = (EmailAddresses) i.next();
            System.out.println(tmpemail.getOwnerAddress());
        }
        System.out.println("the email is CC to:");
        LinkedList<AddressInterface> tmplink2 = msg.getCC();
        for (Iterator i = tmplink2.iterator(); i.hasNext();) {
            EmailAddresses tmpemail = (EmailAddresses) i.next();
            System.out.println(tmpemail.getOwnerAddress());
        }

    }

    public static void ComposeMessage() {
        System.out.println("Please enter the message Title:");
        String title = Globals.scanner.nextLine();
        System.out.println("Please enter the message Description:");
        String Description = Globals.scanner.nextLine();
        System.out.println("Send to:");
        LinkedList<AddressInterface> SentTo = addEmails();
        System.out.println("CC:");
        LinkedList<AddressInterface> CC = addEmails();
        System.out.println("end");
        Messages message = new Messages(title, Description, SentTo, CC);
        Globals.allMessages.add(message);
        Globals.CurrentEmail.setSent(message);
    }

    public static LinkedList<AddressInterface> addEmails() {
        LinkedList<AddressInterface> tmplink = new LinkedList<AddressInterface>();
        while (true) {
            System.out.println("Please enter the emails you want to add one at a time:( or 0 to exit the loop)");
            String email = Globals.scanner.next();
            if (email.equals("0"))
                break;
            EmailAddresses emailAddress = new EmailAddresses(email);
            boolean x = true;
            for (Iterator i = Globals.allEmails.iterator(); i.hasNext();) {
                EmailAddresses tmp = (EmailAddresses) i.next();
                if (email.equals((tmp.getOwnerAddress()))) {
                    tmplink.add(tmp);
                    x = false;
                }
            }
            if (x) {
                tmplink.add(emailAddress);
                Globals.allEmails.add(emailAddress);
            }
        }
        return tmplink;
    }

}
