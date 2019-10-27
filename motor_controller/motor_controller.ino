// drive controller: L293D
#define left_0 6 // pin 14 L293D
#define left_1 7 // pin 11 L293D
#define right_0 4 // pin 3 L293D
#define right_1 5 // pin 6 L293D

#define delay_time 500
#define cry_time 1000

// ultrasound sensor
#define echo 8
#define trig 9
#define thres_dist 15

int m_dir; // master direction: 0 stop, 1 recto, 2 left, 3 right, 4 backwards
long duration;
int distance;
int mode; // 0 run, 1 obstacle, 2 crazy

void setup(){
 
    //Set pins as outputs
    pinMode(left_0, OUTPUT);
    pinMode(left_1, OUTPUT);
    pinMode(right_0, OUTPUT);
    pinMode(right_1, OUTPUT);

    pinMode(trig, OUTPUT);
    pinMode(echo, INPUT);
    Serial.begin(9600);
    m_dir = 1;   
    mode = 1;
  
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

    // ultrasound code moment
    // Clears the trigPin
    digitalWrite(trig, LOW);
    delayMicroseconds(2);
    // Sets the trigPin on HIGH state for 10 micro seconds
    digitalWrite(trig, HIGH);
    delayMicroseconds(10);
    digitalWrite(trig, LOW);
    // Reads the echoPin, returns the sound wave travel time in microseconds
    duration = pulseIn(echo, HIGH);
    // Calculating the distance
    distance= duration*0.034/2;

    if (distance < thres_dist) {
        mode = 1;
        m_dir = 0;
        delay(cry_time);
        delay(cry_time);
        delay(cry_time);
        delay(cry_time);
        delay(cry_time);
    }
    else mode = 0;

        // debug
        Serial.print("distance: ");
        Serial.println(distance, DEC);

    if (Serial.available() > 0) {
        int new_dir = Serial.read() - '0';
        if (mode == 0 && new_dir >= 0) m_dir = new_dir; 
        else if (mode == 1 && new_dir == 4) m_dir = new_dir;
    }
}
