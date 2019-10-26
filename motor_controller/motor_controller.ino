// drive controller: L293D
#define left_0 6 // pin 14 L293D
#define left_1 7 // pin 11 L293D
#define right_0 4 // pin 3 L293D
#define right_1 5 // pin 6 L293D
#define delay_time 500

int m_dir; // master direction: 0 stop, 1 recto, 2 left, 3 right, 4 backwards

void setup(){
 
    //Set pins as outputs
    pinMode(left_0, OUTPUT);
    pinMode(left_1, OUTPUT);
    pinMode(right_0, OUTPUT);
    pinMode(right_1, OUTPUT);
    Serial.begin(9600);
    m_dir = 0;   
  
}

void direction (int a, int b, int c, int d) {
    digitalWrite(left_0, a);
    digitalWrite(left_1, b);
    digitalWrite(right_0, c);
    digitalWrite(right_1, d);
    delay(delay_time); 
}

void loop(){
  
      //Motor Control - Motor A: motorPin1,motorpin2 & Motor B: motorpin3,motorpin4
    switch (m_dir) {
        case 1: // forward
            direction (LOW, HIGH, LOW, HIGH);
            break;
        case 2: // left
            direction (LOW, LOW, LOW, HIGH);
            break;
        case 3: // right
            direction (LOW, HIGH, LOW, LOW);
            break;
        case 4: // backwards
            direction (HIGH, LOW, HIGH, LOW);
            break;
        default:
            direction (LOW, LOW, LOW, LOW);
    }
    // back
/*
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
*/
    if (Serial.available() > 0) {
        int new_dir = Serial.read() - '0';
        if (new_dir >= 0) m_dir = new_dir; 
        // debug
        Serial.print("direction: ");
        Serial.println(m_dir, DEC);
    }
}
