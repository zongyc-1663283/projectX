import React, {Component} from 'react';
import { StyleSheet, Text, View } from 'react-native';
import { Button, TouchableOpacity } from 'react-native';
import t from 'tcomb-form-native';
import axios from 'axios';

const User = t.struct({
  name: t.String,
  start: t.String,
  destination: t.String,
  phone: t.String
});
const Form = t.form.Form;

export default class App extends React.Component {
  geocode = () => {
    const value = this._form.getValue(); // use that ref to get the form value
    var start, destination
    if (value == null) {
      alert("form cannot be empty, do you want to walk home by yourself?")
      return
    }
    axios.get('https://maps.googleapis.com/maps/api/geocode/json',{
      params:{
        address: value.destination,
        key:'AIzaSyCmQvqJq6n6-oqpy_9BHEA2MHiYfLTQ4L4'
      }
    })
    .then(function(response){
      destDetail = response.data.results[0].geometry
      console.log(destDetail)
    });
  }

  render() {
    return (
      <View style={styles.container}>
        <Text style={styles.header}>You never walk alone</Text>
        <Form
          ref={c => this._form = c} // assign a ref
          style={styles.form}
          type={User} 
        /> {/* Notice the addition of the Form component */}
        <TouchableOpacity
          style={styles.customBtnBG}
          onPress={this.geocode}  >
          <Text style={styles.customBtnText}>Walk me</Text>
        </TouchableOpacity>
      </View>
    );
  }
}

const styles = StyleSheet.create({
  container: {
    justifyContent: 'center',
    marginTop: 50,
    padding: 20,
    backgroundColor: '#ffffff',
  },
  header: {
    justifyContent: 'center',
    backgroundColor: '#00c07f',
    fontSize:30,
  },
  customBtnText: {
        fontSize: 47,
        fontWeight: '400',
        color: "#fff",
    },

  /* Here, style the background of your button */
  customBtnBG: {
    backgroundColor: "#007aff",
    paddingHorizontal: 30,
    paddingVertical: 5,
    borderRadius: 30
  }
});