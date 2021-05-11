import java.util.*;

public class TestMessages {
public class TestMessagesClass {
  public void TestgetSentTo() {
    //
    AddressInterface fakeAddress = new AddressInterface() {
      public void method() {
      }
    };
    LinkedList<AddressInterface> fakelist = new LinkedList<AddressInterface>();
    fakelist.add(fakeAddress);

    Messages m = new Messages("title", "Discription", fakelist, fakelist);
    m.getSentTo();
  }

  public void TestgetCC() {
    //
  }

  public void TestgetTitle() {
    //
  }

  public void TestgetDescription() {
    //
  }

  public void TestMessage() {
    //
  }
}
