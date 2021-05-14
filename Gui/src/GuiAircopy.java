import java.awt.Color;
import java.awt.Graphics;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;

import javax.swing.JButton;
import javax.swing.JComboBox;
import javax.swing.JFrame;
import javax.swing.JLabel;
import javax.swing.JPanel;
public class GuiAircopy extends JPanel implements ActionListener{
    String SelectedFlight;
    JButton[] ButtonsA = new JButton[5];
    JButton[] ButtonsB = new JButton[5];
    JButton[] ButtonsC = new JButton[5];
    JButton[] ButtonsD = new JButton[5];
    JComboBox<String> Flights;
    JLabel label;
    String[] FlightNames = {"AUI22 London to Melbourne","AUI33 Berlin to Texas","AUI44 Moscow to Madrid"};
    public void SetButtons(JButton[] Buttons,String s,int strtX,int strtY){
        for (int i = 0; i < Buttons.length; i++) {
            Buttons[i] = new JButton((i+1)+s);
            Buttons[i].setBackground(Color.GREEN);
            Buttons[i].setBounds(strtX+150*i, strtY, 60, 60);
            Buttons[i].addActionListener(this);
            this.add(Buttons[i]);
        }
    }
    public GuiAircopy(){
        setLayout(null);
        Flights = new JComboBox<String>(FlightNames);
        Flights.setMaximumRowCount(3);
        Flights.setBounds(700, 0, 225, 35);
        Flights.setBackground(Color.WHITE);
        Flights.addActionListener(new ActionListener(){
            @Override
            public void actionPerformed(ActionEvent e){
                if(((String)Flights.getSelectedItem()).equals("AUI22 London to Melbourne")){
                    ButtonsA[1].setBackground(Color.BLUE);
                }else if(((String)Flights.getSelectedItem()).equals("AUI33 Berlin to Texas")){
                    ButtonsA[2].setBackground(Color.BLUE);
                }else{
                    ButtonsA[3].setBackground(Color.BLUE);
                }
            }
        });
        add(Flights);
        label = new JLabel("Flights:");
        label.setBounds(630, 0, 200, 35);
        add(label);
        SetButtons(ButtonsA,"A",500,525);
        SetButtons(ButtonsB,"B",500,450);
        SetButtons(ButtonsC,"C",500,325);
        SetButtons(ButtonsD,"D",500,250);
        
    }
    @Override
    public void paintComponent(Graphics g){
        super.paintComponent(g);
        g.drawArc(50, 200, 1000, 400, 90, 180);
        g.drawLine(550, 200, 1400, 200);
        g.drawLine(550, 600, 1400, 600);
        g.drawLine(600, 200, 800, 100);
        g.drawLine(850, 200, 950, 100);
        g.drawLine(550, 600, 800, 700);
        g.drawLine(850, 600, 950, 700);
    }
    public static void main(String[] args) {
        JFrame frame = new JFrame();
        frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        GuiAircopy arc = new GuiAircopy();
        frame.add(arc);
        frame.setSize(1400,925);
        frame.setVisible(true);
    }
    @Override
    public void actionPerformed(ActionEvent e) {
        if(((JButton)e.getSource()).getBackground() == Color.GREEN){
            ((JButton)e.getSource()).setBackground(Color.RED);
        }else {
            ((JButton)e.getSource()).setBackground(Color.GREEN);
        }
        
    }
   
}

