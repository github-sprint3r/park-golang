function CarOutCtrl($scope, $http) {
  $scope.provinces = ["กระบี่","กรุงเทพมหานคร","กาญจนบุรี","กาฬสินธุ์","กำแพงเพชร","ขอนแก่น","จันทบุรี","ฉะเชิงเทรา" ,"ชลบุรี","ชัยนาท","ชัยภูมิ","ชุมพร","เชียงราย","เชียงใหม่","ตรัง","ตราด","ตาก","นครนายก","นครปฐม","นครพนม","นครราชสีมา" ,"นครศรีธรรมราช","นครสวรรค์","นนทบุรี","นราธิวาส","น่าน","บุรีรัมย์","ปทุมธานี","ประจวบคีรีขันธ์","ปราจีนบุรี","ปัตตานี" ,"พะเยา","พังงา","พัทลุง","พิจิตร","พิษณุโลก","เพชรบุรี","เพชรบูรณ์","แพร่","ภูเก็ต","มหาสารคาม","มุกดาหาร","แม่ฮ่องสอน" ,"ยโสธร","ยะลา","ร้อยเอ็ด","ระนอง","ระยอง","ราชบุรี","ลพบุรี","ลำปาง","ลำพูน","เลย","ศรีสะเกษ","สกลนคร","สงขลา" ,"สตูล","สมุทรปราการ","สมุทรสงคราม","สมุทรสาคร","สระแก้ว","สระบุรี","สิงห์บุรี","สุโขทัย","สุพรรณบุรี","สุราษฎร์ธานี" ,"สุรินทร์","หนองคาย","หนองบัวลำภู","อยุธยา","อ่างทอง","อำนาจเจริญ","อุดรธานี","อุตรดิตถ์","อุทัยธานี","อุบลราชธานี"];
  $scope.selected_province = $scope.provinces[1];
  $scope.data = {
	  "license_number": "",
	  "province": "",
	  "in_time": "วันจันทร์ที่ 5 มกราคม 2557 เวลา 09:30 น.",
	  "out_time": "วันจันทร์ที่ 5 มกราคม 2557 เวลา 12:40 น.",
	  "fee": 0,
  };
  $scope.license_number = "";
  
  $scope.url = "http://www.google.co.th/?gws_rd=cr,ssl&ei=Br63U5OZF9OKuATqt4DIBQ";
  $scope.fetch = function() {
	  $scope.hide_result();
	  $scope.loading();
	  var data = {"license_number": $scope.license_number, "province": $scope.provinces[$scope.selected_province]};
	  $scope.data.license_number = $scope.license_number;
	  $scope.data.province = $scope.selected_province;
	  
	  $http.post($scope.url, data, {withCredentials: true, headers: {'Access-Control-Allow-Origin': '*', "Access-Control-Allow-Origin": "Origin, X-Requested-With, Content-Type, Accept"}}).
	  //$http({
		//  withCredentials: true,
		//  method: 'POST',
		//  url: $scope.url,
		//  data: $scope.data,
		//  headers: {"Access-Control-Allow-Origin": "*", "Access-Control-Allow-Methods": "POST, GET, OPTIONS, PUT", "Access-Control-Allow-Origin": "Origin, X-Requested-With, Content-Type, Accept"}
	  //}).
	    success(function(data, status, headers, config) {
	      // this callback will be called asynchronously
	      // when the response is available
		  $scope.load_finished();
	    }).
	    error(function(data, status, headers, config) {
	      // called asynchronously if an error occurs
	      // or server returns response with an error status.
		  $scope.load_failed();
	    });
	//.then(function(response) {
    //$scope.load_finished();
	//});
  }
  
  $scope._loading = false;
  $scope.loading = function() {
	  $scope._loading = true;
  }
  $scope.stop_loading = function() {
	  $scope._loading = false;
  }
  $scope.is_loading = function() {
	  return $scope._loading;
  }
  
  $scope._result_loaded = false;
  $scope.is_result_loaded = function() {
	  return $scope._result_loaded;
  }
  $scope.load_finished = function() {
	  $scope.stop_loading();
	  $scope._result_loaded = true;
  }
  $scope.load_failed = function() {
	  $scope.stop_loading();
	  $scope._result_loaded = false;  	
  }
  $scope.hide_result = function() {
	  $scope._result_loaded = false;  	
  }
}
