import java.util.*;

public class TestMessagesClass {
  public void TestgetSentTo() {
    //
    AddressInterface fakeAddress = new AddressInterface() {
      public void Print() {
      }
    };
    LinkedList<AddressInterface> fakelist = new LinkedList<AddressInterface>();
    fakelist.add(fakeAddress);

    Messages m = new Messages("title", "Discription", fakelist, fakelist);
    m.getSentTo();
  }

  public void TestgetCC() {
    //
    AddressInterface fakeAddress = new AddressInterface() {
      public void Print() {
      }
    };
    LinkedList<AddressInterface> fakelist = new LinkedList<AddressInterface>();
    fakelist.add(fakeAddress);

    Messages m = new Messages("title", "Discription", fakelist, fakelist);
    m.getCC();
  }

  public void TestgetTitle() {
    //
    AddressInterface fakeAddress = new AddressInterface() {
      public void Print() {
      }
    };
    LinkedList<AddressInterface> fakelist = new LinkedList<AddressInterface>();
    fakelist.add(fakeAddress);

    Messages m = new Messages("title", "Discription", fakelist, fakelist);
    m.getTitle();
  }

  public void TestgetDescription() {
    //
    AddressInterface fakeAddress = new AddressInterface() {
      public void Print() {
      }
    };
    LinkedList<AddressInterface> fakelist = new LinkedList<AddressInterface>();
    fakelist.add(fakeAddress);

    Messages m = new Messages("title", "Discription", fakelist, fakelist);
    m.getDescription();
  }

  public void TestMessages() {
    //
    AddressInterface fakeAddress = new AddressInterface() {
      public void Print() {
      }
    };
    LinkedList<AddressInterface> fakelist = new LinkedList<AddressInterface>();
    fakelist.add(fakeAddress);

    Messages m = new Messages("title", "Discription", fakelist, fakelist);
  }
}
