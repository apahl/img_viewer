package main

var indexHTML = []byte(`
<!doctype html>
<html>

<head>
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <style>
        body {
            text-align: left;
            color: black;
            background-color: #FFFFFF;
            font-family: freesans, arial, verdana, sans-serif;
        }

        button {
            color: black;
        }

        table {
            border-collapse: collapse;
            border-width: thin;
            border-style: solid;
            border: none;
            background-color: #FFFFFF;
            text-align: left;
        }

        th {
            border-collapse: collapse;
            border-width: thin;
            border-style: solid;
            border-color: black;
            background-color: #94CAEF;
            text-align: center;
            font-weight: bold;
            padding: 5px;
        }

        td {
            border-collapse: collapse;
            border-width: thin;
            border-style: solid;
            border-color: black;
            text-align: center;
            padding: 5px;
        }

        td.separate {
            border-collapse: separate;
            border-width: thin;
            border-style: solid;
            border-color: black;
            text-align: center;
            padding: 5px;
        }

        table.noborder {
            border-collapse: separate;
            border-width: thin;
            border-style: solid;
            border: none;
            background-color: #FFFFFF;
            text-align: left;
        }

        th.noborder {
            border-collapse: separate;
            border-width: thin;
            border-style: solid;
            border-color: white;
            background-color: #94CAEF;
            text-align: left;
            font-weight: bold;
            padding: 5px;
        }

        td.noborder {
            border-collapse: separate;
            border-width: thin;
            border-style: solid;
            text-align: left;
            border-color: white;
            padding: 5px;
        }

        th.left {
            text-align: left;
        }

        td.left {
            text-align: left;
        }
    </style>
</head>

<body>
    Base Path:
    <input id="basepath" type="text" value="" size=80 />
    <button onclick="external.invoke('config:')">
        Open previous set of plates
    </button>
    <br> Slot 1:
    <input id="plate-id1" type="text" value="" />
    <input id="well-id1" type="text" value="" size=4 size=4 />
    <button onclick="external.invoke('next:' + '1,' + document.getElementById('basepath').value + ',' + document.getElementById('plate-id1').value + ',' + document.getElementById('well-id1').value + ',-24,' + document.getElementById('autoload').checked)">
        &uarr;
    </button>
    <button onclick="external.invoke('next:' + '1,' + document.getElementById('basepath').value + ',' + document.getElementById('plate-id1').value + ',' + document.getElementById('well-id1').value + ',24,' + document.getElementById('autoload').checked)">
        &darr;
    </button>
    <button onclick="external.invoke('next:' + '1,' + document.getElementById('basepath').value + ',' + document.getElementById('plate-id1').value + ',' + document.getElementById('well-id1').value + ',-1,' + document.getElementById('autoload').checked)">
        &larr;
    </button>
    <button onclick="external.invoke('next:' + '1,' + document.getElementById('basepath').value + ',' + document.getElementById('plate-id1').value + ',' + document.getElementById('well-id1').value + ',1,' + document.getElementById('autoload').checked)">
        &rarr;
    </button>
    <button onclick="external.invoke('load:' + '1,' + document.getElementById('basepath').value + ',' + document.getElementById('plate-id1').value + ',' + document.getElementById('well-id1').value)">
        Load
    </button>

    &nbsp;&nbsp;&nbsp;&nbsp; Slot 4:
    <input id="plate-id4" type="text" value="" />
    <input id="well-id4" type="text" value="" size=4 />
    <button onclick="external.invoke('next:' + '4,' + document.getElementById('basepath').value + ',' + document.getElementById('plate-id4').value + ',' + document.getElementById('well-id4').value + ',-24,' + document.getElementById('autoload').checked)">
        &uarr;
    </button>
    <button onclick="external.invoke('next:' + '4,' + document.getElementById('basepath').value + ',' + document.getElementById('plate-id4').value + ',' + document.getElementById('well-id4').value + ',24,' + document.getElementById('autoload').checked)">
        &darr;
    </button>
    <button onclick="external.invoke('next:' + '4,' + document.getElementById('basepath').value + ',' + document.getElementById('plate-id4').value + ',' + document.getElementById('well-id4').value + ',-1,' + document.getElementById('autoload').checked)">
        &larr;
    </button>
    <button onclick="external.invoke('next:' + '4,' + document.getElementById('basepath').value + ',' + document.getElementById('plate-id4').value + ',' + document.getElementById('well-id4').value + ',1,' + document.getElementById('autoload').checked)">
        &rarr;
    </button>
    <button onclick="external.invoke('load:' + '4,' + document.getElementById('basepath').value + ',' + document.getElementById('plate-id4').value + ',' + document.getElementById('well-id4').value)">
        Load
    </button>
    &nbsp;&nbsp;&nbsp;&nbsp;All:&nbsp;
    <button onclick="external.invoke('sync:' + document.getElementById('basepath').value + ',' + document.getElementById('plate-id1').value + ',' +
                                  document.getElementById('plate-id2').value + ',' +
                                  document.getElementById('plate-id3').value + ',' +
                                  document.getElementById('plate-id4').value + ',' +
                                  document.getElementById('plate-id5').value + ',' +
                                  document.getElementById('plate-id6').value + ',' +
                                  document.getElementById('well-id1').value + ',' +
                                  document.getElementById('well-id2').value + ',' +
                                  document.getElementById('well-id3').value + ',' +
                                  document.getElementById('well-id4').value + ',' +
                                  document.getElementById('well-id5').value + ',' +
                                  document.getElementById('well-id6').value + ',-24')">
        &uarr;
    </button>
    <button onclick="external.invoke('sync:' + document.getElementById('basepath').value + ',' + document.getElementById('plate-id1').value + ',' +
                                  document.getElementById('plate-id2').value + ',' +
                                  document.getElementById('plate-id3').value + ',' +
                                  document.getElementById('plate-id4').value + ',' +
                                  document.getElementById('plate-id5').value + ',' +
                                  document.getElementById('plate-id6').value + ',' +
                                  document.getElementById('well-id1').value + ',' +
                                  document.getElementById('well-id2').value + ',' +
                                  document.getElementById('well-id3').value + ',' +
                                  document.getElementById('well-id4').value + ',' +
                                  document.getElementById('well-id5').value + ',' +
                                  document.getElementById('well-id6').value + ',24')">
        &darr;
    </button>
    <button onclick="external.invoke('sync:' + document.getElementById('basepath').value + ',' + document.getElementById('plate-id1').value + ',' +
                                  document.getElementById('plate-id2').value + ',' +
                                  document.getElementById('plate-id3').value + ',' +
                                  document.getElementById('plate-id4').value + ',' +
                                  document.getElementById('plate-id5').value + ',' +
                                  document.getElementById('plate-id6').value + ',' +
                                  document.getElementById('well-id1').value + ',' +
                                  document.getElementById('well-id2').value + ',' +
                                  document.getElementById('well-id3').value + ',' +
                                  document.getElementById('well-id4').value + ',' +
                                  document.getElementById('well-id5').value + ',' +
                                  document.getElementById('well-id6').value + ',-1')">
        &larr;
    </button>
    <button onclick="external.invoke('sync:' + document.getElementById('basepath').value + ',' + document.getElementById('plate-id1').value + ',' +
                                  document.getElementById('plate-id2').value + ',' +
                                  document.getElementById('plate-id3').value + ',' +
                                  document.getElementById('plate-id4').value + ',' +
                                  document.getElementById('plate-id5').value + ',' +
                                  document.getElementById('plate-id6').value + ',' +
                                  document.getElementById('well-id1').value + ',' +
                                  document.getElementById('well-id2').value + ',' +
                                  document.getElementById('well-id3').value + ',' +
                                  document.getElementById('well-id4').value + ',' +
                                  document.getElementById('well-id5').value + ',' +
                                  document.getElementById('well-id6').value + ',1')">
        &rarr;
    </button>

    <br> Slot 2:
    <input id="plate-id2" type="text" value="" />
    <input id="well-id2" type="text" value="" size=4 />
    <button onclick="external.invoke('next:' + '2,' + document.getElementById('basepath').value + ',' + document.getElementById('plate-id2').value + ',' + document.getElementById('well-id2').value + ',-24,' + document.getElementById('autoload').checked)">
        &uarr;
    </button>
    <button onclick="external.invoke('next:' + '2,' + document.getElementById('basepath').value + ',' + document.getElementById('plate-id2').value + ',' + document.getElementById('well-id2').value + ',24,' + document.getElementById('autoload').checked)">
        &darr;
    </button>
    <button onclick="external.invoke('next:' + '2,' + document.getElementById('basepath').value + ',' + document.getElementById('plate-id2').value + ',' + document.getElementById('well-id2').value + ',-1,' + document.getElementById('autoload').checked)">
        &larr;
    </button>
    <button onclick="external.invoke('next:' + '2,' + document.getElementById('basepath').value + ',' + document.getElementById('plate-id2').value + ',' + document.getElementById('well-id2').value + ',1,' + document.getElementById('autoload').checked)">
        &rarr;
    </button>
    <button onclick="external.invoke('load:' + '2,' + document.getElementById('basepath').value + ',' + document.getElementById('plate-id2').value + ',' + document.getElementById('well-id2').value)">
        Load
    </button>

    &nbsp;&nbsp;&nbsp;&nbsp; Slot 5:
    <input id="plate-id5" type="text" value="" />
    <input id="well-id5" type="text" value="" size=4 />
    <button onclick="external.invoke('next:' + '5,' + document.getElementById('basepath').value + ',' + document.getElementById('plate-id5').value + ',' + document.getElementById('well-id5').value + ',-24,' + document.getElementById('autoload').checked)">
        &uarr;
    </button>
    <button onclick="external.invoke('next:' + '5,' + document.getElementById('basepath').value + ',' + document.getElementById('plate-id5').value + ',' + document.getElementById('well-id5').value + ',24,' + document.getElementById('autoload').checked)">
        &darr;
    </button>
    <button onclick="external.invoke('next:' + '5,' + document.getElementById('basepath').value + ',' + document.getElementById('plate-id5').value + ',' + document.getElementById('well-id5').value + ',-1,' + document.getElementById('autoload').checked)">
        &larr;
    </button>
    <button onclick="external.invoke('next:' + '5,' + document.getElementById('basepath').value + ',' + document.getElementById('plate-id5').value + ',' + document.getElementById('well-id5').value + ',1,' + document.getElementById('autoload').checked)">
        &rarr;
    </button>
    <button onclick="external.invoke('load:' + '5,' + document.getElementById('basepath').value + ',' + document.getElementById('plate-id5').value + ',' + document.getElementById('well-id5').value)">
        Load
    </button>

    <br> Slot 3:
    <input id="plate-id3" type="text" value="" />
    <input id="well-id3" type="text" value="" size=4 />
    <button onclick="external.invoke('next:' + '3,' + document.getElementById('basepath').value + ',' + document.getElementById('plate-id3').value + ',' + document.getElementById('well-id3').value + ',-24,' + document.getElementById('autoload').checked)">
        &uarr;
    </button>
    <button onclick="external.invoke('next:' + '3,' + document.getElementById('basepath').value + ',' + document.getElementById('plate-id3').value + ',' + document.getElementById('well-id3').value + ',24,' + document.getElementById('autoload').checked)">
        &darr;
    </button>
    <button onclick="external.invoke('next:' + '3,' + document.getElementById('basepath').value + ',' + document.getElementById('plate-id3').value + ',' + document.getElementById('well-id3').value + ',-1,' + document.getElementById('autoload').checked)">
        &larr;
    </button>
    <button onclick="external.invoke('next:' + '3,' + document.getElementById('basepath').value + ',' + document.getElementById('plate-id3').value + ',' + document.getElementById('well-id3').value + ',1,' + document.getElementById('autoload').checked)">
        &rarr;
    </button>
    <button onclick="external.invoke('load:' + '3,' + document.getElementById('basepath').value + ',' + document.getElementById('plate-id3').value + ',' + document.getElementById('well-id3').value)">
        Load
    </button>

    &nbsp;&nbsp;&nbsp;&nbsp; Slot 6:
    <input id="plate-id6" type="text" value="" />
    <input id="well-id6" type="text" value="" size=4 />
    <button onclick="external.invoke('next:' + '6,' + document.getElementById('basepath').value + ',' + document.getElementById('plate-id6').value + ',' + document.getElementById('well-id6').value + ',-24,' + document.getElementById('autoload').checked)">
        &uarr;
    </button>
    <button onclick="external.invoke('next:' + '6,' + document.getElementById('basepath').value + ',' + document.getElementById('plate-id6').value + ',' + document.getElementById('well-id6').value + ',24,' + document.getElementById('autoload').checked)">
        &darr;
    </button>
    <button onclick="external.invoke('next:' + '6,' + document.getElementById('basepath').value + ',' + document.getElementById('plate-id6').value + ',' + document.getElementById('well-id6').value + ',-1,' + document.getElementById('autoload').checked)">
        &larr;
    </button>
    <button onclick="external.invoke('next:' + '6,' + document.getElementById('basepath').value + ',' + document.getElementById('plate-id6').value + ',' + document.getElementById('well-id6').value + ',1,' + document.getElementById('autoload').checked)">
        &rarr;
    </button>
    <button onclick="external.invoke('load:' + '6,' + document.getElementById('basepath').value + ',' + document.getElementById('plate-id6').value + ',' + document.getElementById('well-id6').value)">
        Load
    </button>
    &nbsp;&nbsp;
    <input id="autoload" type="checkbox" name="autoload"> autoload
    <br>
    <br>
    <table>
        <tr>
            <th>Slot</th>
            <th>Well</th>
            <th>W1</th>
            <th>W2</th>
            <th>W3</th>
            <th>W4</th>
            <th>W5</th>
        </tr>
        <tr>
            <td>
                <b>1</b>
            </td>
            <td>
                <b id="well1"></b>
            </td>
            <td id="w1_1"></td>
            <td id="w1_2"></td>
            <td id="w1_3"></td>
            <td id="w1_4"></td>
            <td id="w1_5"></td>
        </tr>
        <tr>
            <td>
                <b>2</b>
            </td>
            <td>
                <b id="well2"></b>
            </td>
            <td id="w2_1"></td>
            <td id="w2_2"></td>
            <td id="w2_3"></td>
            <td id="w2_4"></td>
            <td id="w2_5"></td>
        </tr>
        <tr>
            <td>
                <b>3</b>
            </td>
            <td>
                <b id="well3"></b>
            </td>
            <td id="w3_1"></td>
            <td id="w3_2"></td>
            <td id="w3_3"></td>
            <td id="w3_4"></td>
            <td id="w3_5"></td>
        </tr>
        <tr>
            <td>
                <b>4</b>
            </td>
            <td>
                <b id="well4"></b>
            </td>
            <td id="w4_1"></td>
            <td id="w4_2"></td>
            <td id="w4_3"></td>
            <td id="w4_4"></td>
            <td id="w4_5"></td>
        </tr>
        <tr>
            <td>
                <b>5</b>
            </td>
            <td>
                <b id="well5"></b>
            </td>
            <td id="w5_1"></td>
            <td id="w5_2"></td>
            <td id="w5_3"></td>
            <td id="w5_4"></td>
            <td id="w5_5"></td>
        </tr>
        <tr>
            <td>
                <b>6</b>
            </td>
            <td>
                <b id="well6"></b>
            </td>
            <td id="w6_1"></td>
            <td id="w6_2"></td>
            <td id="w6_3"></td>
            <td id="w6_4"></td>
            <td id="w6_5"></td>
        </tr>
    </table>
    <br>
    <br>
    <p>
        <small>COMAS 2018, License: MIT; written in Go (https://golang.org/).
            <br> source code available on Github: https://github.com/apahl/img_viewer</small>
    </p>
</body>

</html>`)
