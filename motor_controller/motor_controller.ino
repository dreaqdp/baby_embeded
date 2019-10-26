// drive controller: L293D
#define left_0 6 // pin 14 L293D
#define left_1 7 // pin 11 L293D
#define right_0 4 // pin 3 L293D
#define right_1 5 // pin 6 L293D


int m_dir; // master direction: 0 stop, 1 recto, 2 left, 3 right, 4 backwards

void setup(){
 
    //Set pins as outputs
    pinMode(left_0, OUTPUT);
    pinMode(left_1, OUTPUT);
    pinMode(right_0, OUTPUT);
    pinMode(right_1, OUTPUT);
    Serial.begin(9600);
    m_dir = 1;   
  
}


void loop(){
  
      //Motor Control - Motor A: motorPin1,motorpin2 & Motor B: motorpin3,motorpin4

    // back
    digitalWrite(left_0, HIGH);
    digitalWrite(left_1, LOW);
    digitalWrite(right_0, HIGH);
    digitalWrite(right_1, LOW);
    delay(2000); 

    // right
    digitalWrite(left_0, LOW);
    digitalWrite(left_1, HIGH);
    digitalWrite(right_0, LOW);
    digitalWrite(right_1, LOW);
    delay(2000);

    // left
    digitalWrite(left_0, LOW);
    digitalWrite(left_1, LOW);
    digitalWrite(right_0, LOW);
    digitalWrite(right_1, HIGH);
    delay(2000);

    // forwards
    digitalWrite(left_0, LOW);
    digitalWrite(left_1, HIGH);
    digitalWrite(right_0, LOW);
    digitalWrite(right_1, HIGH);
    delay(2000);

    // stop
    digitalWrite(left_0, LOW);
    digitalWrite(left_1, LOW);
    digitalWrite(right_0, LOW);
    digitalWrite(right_1, LOW);
    delay(2000);
    /*
    //This code will turn Motor B clockwise for 2 sec.
    digitalWrite(motorPin1, LOW);
    digitalWrite(motorPin2, LOW);
    digitalWrite(motorPin3, HIGH);
    digitalWrite(motorPin4, LOW);
    delay(2000); 
    //This code will turn Motor B counter-clockwise for 2 sec.
    digitalWrite(motorPin1, LOW);
    digitalWrite(motorPin2, LOW);
    digitalWrite(motorPin3, LOW);
    digitalWrite(motorPin4, HIGH);
    delay(2000);    
    
    //And this code will stop motors
    digitalWrite(motorPin1, LOW);
    digitalWrite(motorPin2, LOW);
    digitalWrite(motorPin3, LOW);
    digitalWrite(motorPin4, LOW);*/
}
