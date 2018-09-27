var airdbApp = angular.module('airdbApp', ['ngRoute']);
airdbApp.config(function($routeProvider){
    $routeProvider
        .when('/', {
            templateUrl: 'ui/assets/pages/dashboard.html',
            controller: 'mainController'
        })
        .when('/dbs', {
            templateUrl: 'ui/assets/pages/dbs.html',
            controller: 'dbController'
        })
        .when('/dbs/add', {
            templateUrl: 'ui/assets/pages/addDb.html',
            controller: 'addDbController'
        })
        .when('/users', {
            templateUrl: 'ui/assets/pages/users.html',
            controller: 'userController'
        })
});

airdbApp.controller('mainController', ['$scope', '$http', '$location', function($scope, $http, $location) {
    userId = $("#userId").val();
    var editor = ace.edit("editor");
    editor.setTheme("ace/theme/dracula");
    editor.session.setMode("ace/mode/sql");
    $http({ url: 'http://localhost:4000/dbs', method: "POST", headers: {'Content-Type': 'application/json'}, data: { 'userId' : userId }}).then(function(response) {
        var res = response.data
        if(res.code == "100"){
            $scope.Dbs = res.Dbs; 
        }
    });
    $scope.runQuery = function () {
        var query = editor.getValue();
        t = query.substring(query.length - 1);
        if(t != ";"){
            $scope.message = false;
            $scope.error_message = 'Invalid syntax. Please review your syntax to contine';
        }else{
            query = query.slice(0, -1); 
            let Cols = [];
            $http({
                url: 'http://localhost:4000/queries/run',
                method: "POST",
                headers: {'Content-Type': 'application/json'},
                data: { 'dbId' : $scope.dbId, 'userId': userId,  'query': query}
                }).then(function(response) {
                    var res = response.data
                    if(res.code == "100"){
                        dt = res.data;
                        $scope.Data = dt;
                        dt.forEach(e => {
                        Cols = Object.keys(e);
                        });
                        $scope.Columns = Cols
                        $scope.message = true;
                    } else {
                        $scope.message = false;
                        $scope.error_message = res.error;
                    }
                });
        }
    }
    $scope.isEmpty = function (obj) {
        for (var i in obj) if (obj.hasOwnProperty(i)) return false;
        return true;
    };

}]);

airdbApp.controller('dbController', ['$scope','$http', function($scope, $http) {
    userId = $("#userId").val();
    $http({ url: 'http://localhost:4000/dbs', method: "POST", headers: {'Content-Type': 'application/json'}, data: { 'userId' : userId }}).then(function(response) {
        var res = response.data
        let Cols = [];
        if(res.code == "100"){
            $scope.Dbs = res.Dbs;
            $scope.Dbs = Object.keys(res.Dbs)
            .map(function (value, index) {
                return { joinDate: new Date(value), values: res.Dbs[value] }
            });
            dt = res.Dbs;
            dt.forEach(e => {
                Cols = Object.keys(e);
            });
            $scope.Columns = Cols;
        }
    });
}]);


airdbApp.controller('addDbController', ['$scope','$http', '$location', function($scope, $http, $location) {
    userId = $("#userId").val();
    $scope.saveDb = function() {
        name = $("#name").val();
        dbschema = $("#dbschema").val();
        host = $("#host").val();
        username = $("#username").val();
        password = $("#password").val();
        port = $("#port").val();
        $http({
            url: 'http://localhost:4000/dbs/add',
            method: "POST",
            headers: {'Content-Type': 'application/json'},
            data: {'name':name,'dbschema': dbschema,'host':host, 'port': port, 'username': username,'password': password, 'userId': userId}
        }).then(function(response) {
            var res = response.data
            if(res.code == "100"){
                $scope.message = true;
            }else{
                $scope.message = false;
                $scope.error_message = res.error;
            }
        });

    }
}]);


airdbApp.controller('userController', ['$scope','$http', function($scope, $http) {
    $http({ url: 'http://localhost:4000/users', method: "GET", headers: {'Content-Type': 'application/json'}}).then(function(response) {
        var res = response.data
        let Cols = [];
        if(res.code == "100"){
            $scope.users = Object.keys(res.users)
            .map(function (value, index) {
                return { joinDate: new Date(value), values: res.users[value] }
            });
            dt = res.users;
            dt.forEach(e => {
                Cols = Object.keys(e);
            });
            $scope.columns = Cols;
        }
    });
}]);

