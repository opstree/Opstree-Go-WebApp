package webapp

const htmltemplate=`{{ define "Index" }}
{{ template "Header" }}
  {{ template "Menu"  }}
  <h2><strong>Registered Users</strong></h2>
  <table border="1" class="table table-bordered">
	<thead>
	<tr>
	  <th>ID</th>
	  <th>Name</th>
	  <th>View</th>
	  <th>Edit</th>
	  <th>Delete</th>
	</tr>
	 </thead>
	 <tbody>
  {{ range . }}
	<tr>
	  <td>{{ .Id }}</td>
	  <td> {{ .Name }} </td>
	  <td><a href="/show?id={{ .Id }}">View</a></td>
	  <td><a href="/edit?id={{ .Id }}">Edit</a></td>
	  <td><a href="/delete?id={{ .Id }}">Delete</a></td>
	</tr>
  {{ end }}
	 </tbody>
  </table>
{{ template "Footer" }}
{{ end }}

{{ define "Header" }}
<!DOCTYPE html>
<html lang="en-US">
    <head>
        <title>OpsTree Golang Sample Crud Application</title>
        <meta charset="UTF-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css">
        <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.4.1/jquery.min.js"></script>
        <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.7/umd/popper.min.js"></script>
        <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js"></script>
        <style>
          img {
            display: block;
            margin-left: auto;
            margin-right: auto;
          }
        </style>
    </head>
    <body>
    <div class="form-row">
    <div class="container">
    <br></br>
    <h2><strong><center>Opstree Golang Sample Crud Application</center></strong></h2>
    <img src="https://res.cloudinary.com/practicaldev/image/fetch/s--3rFO85cD--/c_limit%2Cf_auto%2Cfl_progressive%2Cq_66%2Cw_880/https://thepracticaldev.s3.amazonaws.com/i/bkv3xbjb74epempcjone.gif" style="width:256px;height:256px;text-align:center;">
{{ end }}

{{ define "Footer" }}
</div>
    </body>
</html>
{{ end }}

{{ define "Menu" }}
<br></br>
<a href="/"><strong>HOME</strong></a> | 
<a href="/new"><strong>NEW</strong></a>
<br></br>
{{ end }}

{{ define "Show" }}
  {{ template "Header" }}
    {{ template "Menu"  }}
    <h2><strong>Registeration Number:- {{ .Id }}</strong></h2>
    <table border="1" class="table table-bordered">
    <thead>
    <tr>
      <th>Name</th>
      <th>Email</th>
      <th>Joining Date</th>
      <th>City</th>
    </tr>
    </thead>
    <tbody>
    <tr>
      <td>{{ .Name }}</td>
      <td>{{ .Email }}</td>
      <td>{{ .Date }}</td>
      <td>{{ .City }}</td>
    </tr>
    </tbody>
    </table>
  {{ template "Footer" }}
{{ end }}

{{ define "New" }}
  {{ template "Header" }}
    {{ template "Menu" }}  
    <h2>Create Information</h2>  
    <form method="POST" action="insert">
    <div class="form-group">
      <input type="hidden" name="uid" value="{{ .Id }}" />
    </div>
    <div class="form-group">
      <label for="name">Name:</label>
      <input type="text" name="name" value id="name" class="form-control" placeholder="e.g. Sandeep Rawat">
    </div>
    <div class="form-group">
      <label for="city">City:</label>
      <input type="text" name="city" value id="city" class="form-control" placeholder="e.g. Delhi">
    </div>
    <div class="form-group">
      <label for="email">Email:</label>
      <input type="email" name="email" value id="email" class="form-control" placeholder="e.g. abc@example.com">
    </div>
    <div class="form-group">
      <label for="date">Joining Date:</label>
      <input type="date" name="date" value id="date" class="form-control" placeholder="e.g. 12/12/2012">
    </div>
    <button type="submit" class="btn btn-success">Submit</button>
    </form>
  {{ template "Footer" }}
{{ end }}

{{ define "Edit" }}
  {{ template "Header" }}
    {{ template "Menu" }} 
   <h2><strong>Edit Information for {{ .Name }}</strong></h2>  
    <form method="POST" action="update">
    <div class="form-group">
      <input type="hidden" name="uid" value="{{ .Id }}" />
    </div>
    <div class="form-group">
      <label for="name">Name:</label>
      <input type="text" name="name" value="{{ .Name }}" id="name"  class="form-control"/><br />
    </div>
    <div class="form-group">
      <label for="city">City:</label>
      <input type="text" name="city" value="{{ .City }}" id="city" class="form-control"  /><br />
    </div>
    <div class="form-group">
      <label for="email">Email:</label>
      <input type="email" name="email" value="{{ .Email }}" id="email" class="form-control"  /><br />
    </div>
    <div class="form-group">
      <label for="date">Joining Date:</label>
      <input type="date" name="date" value="{{ .Date }}" id="date" class="form-control"  /><br />
    </div>
      <button type="submit" class="btn btn-success">Submit</button>
    </form><br />    
  {{ template "Footer" }}
{{ end }}`
