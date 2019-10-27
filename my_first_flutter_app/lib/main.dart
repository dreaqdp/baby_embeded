import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';

enum WhyFarther { harder, smarter, selfStarter, tradingCharter }

void main() {
  runApp(new MaterialApp(
    home: new MyApp (),
  ));
}

class MyApp extends StatefulWidget {
  @override
  _MyAppState createState() => new _MyAppState();
}

class _MyAppState extends State<MyApp> {
  var _myColor = [Colors.blue,Colors.lightBlue,Colors.lightBlue,Colors.blue];
  String uri = "http://10.1.123.36:2345/";


  @override
  Widget build(BuildContext context) {
    return new Scaffold(
      appBar: new AppBar(
        title: new Text ("Tap me!"),
        actions: <Widget>[
          IconButton(icon: Icon(Icons.list), onPressed: _optionsbutton)
        ],
      ),
      body: _buildButtons(),
    );
  }

  Widget _buildButtons() {
    var height = MediaQuery.of(context).size.height;
    var width = MediaQuery.of(context).size.width;
    return Container(
      child: Column(
        children: [
          Container(
              height: 0.08 * height,
              child: _postTextfield()
          ),
          Container(
            color: Colors.grey,
            height: 0.43 * height,
            width: MediaQuery.of(context).size.width,
          ),
          _buildButton(0, 1, "forward"),
          Container(
            child: Row(
              children: <Widget>[
                _buildButton(1,2,"left"),
                _buildButton(2,2,"right"),
              ],
            ),
          ),
          _buildButton(3,1,"back"),
        ],
      ),
    );
  }

  void _optionsbutton(){
    Navigator.of(context).push(
        MaterialPageRoute<void>(
            builder: (BuildContext context) {
              return Scaffold(
                  appBar: AppBar(
                      title: Text('Configuration')
                  ),
                  body:
                  Container(
                      child: TextField(
                        onSubmitted: (text) {
                          print(text);
                          uri = text;
                        },
                        decoration: InputDecoration(
                            border: InputBorder.none,
                            hintText: 'Enter a new IP address...'
                        ),
                      ),
                      decoration: BoxDecoration(
                      border: Border.all(color: Colors.black)
                  ),
                  )
              );
            }
            )
    );
  }

  Widget _postTextfield(){
    return TextField(
        onSubmitted: (text) {
          print(text);
          var body = json.encode({"yt": text});
          var headers = {
            'Content-type': 'application/json',
            'Accept': 'application/json'
        };
        (var uri) async { http.post(uri+"play", body: body, headers: headers);}(uri);
        },
      decoration: InputDecoration(
          border: InputBorder.none,
          hintText: 'Enter a youtube music video link...'
      )
    );
  }

  Widget _buildButton(int id, int w, String action) {
    var height = MediaQuery.of(context).size.height;
    return new GestureDetector(
      onTapDown:(TapDownDetails details) { setState(() {
        (var uri, var action) async { http.get(uri+action);}(uri, action);
        _myColor[id] = Colors.orange;
      });
      },
      onTapUp: (TapUpDetails details) { setState(() {
        (var uri, var action) async { http.get(uri+action);}(uri, "stop");
        _myColor[id] = w == 2 ? Colors.lightBlue : Colors.blue;
      });
      },

      child: new Container (
        height: 0.12 * height,
        width: MediaQuery.of(context).size.width/w,
        decoration: BoxDecoration(
            color: _myColor[id],
            border: Border.all(color: Colors.black)
        ),
        child: Center( child: Text(
          action.toUpperCase(),
          style: TextStyle(fontWeight: FontWeight.bold),)
        ),
      ),
    );
  }
}

