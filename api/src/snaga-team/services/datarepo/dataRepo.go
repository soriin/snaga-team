package datarepo
//
// import (
// 	"reflect"
//
// 	"appengine"
// 	"appengine/datastore"
//
// 	"snaga-team/models"
// )
//
// type DataRepo struct {
// 	MyContext appengine.Context
// }
//
// func NewDataRepo(c appengine.Context) *DataRepo {
// 	return &DataRepo{MyContext: c}
// }
//
// func (repo *DataRepo) Put(obj interface{}, kind string, ancestorPath *datastore.Key) (*datastore.Key, error) {
// 		key, err := datastore.Put(repo.MyContext, datastore.NewIncompleteKey(repo.MyContext, kind, ancestorPath), obj)
//
// 		repo.updateId(obj, key)
//
// 		return key, err
// }
//
// func (repo *DataRepo) Query(q *datastore.Query, myType reflect.Type, temp interface{}) (interface{}, error) {
// 	mySlice := reflect.MakeSlice(reflect.SliceOf(myType), 1, 1)
//
// 	for t := q.Run(repo.MyContext); ; {
//     //x := reflect.New(myType)
// 		var test models.Ship
//     key, err := t.Next(&test)
//
//     if err == datastore.Done {
//       break
//     }
//     if err != nil {
//       return mySlice, err
//     }
//     repo.updateId(test, key)
// 		y := reflect.ValueOf(test)
// 		y.Set(test)
//     mySlice = reflect.Append(mySlice, y)
//   }
//
// 	return mySlice, nil
// }
// //
// // func (repo *DataRepo) Get(obj interface{}, kind string, ancestorPath *datastore.Key) (*datastore.Key, error) {
// // 		key, err := datastore.Put(repo.MyContext, datastore.NewIncompleteKey(repo.MyContext, kind, ancestorPath), obj)
// //
// // 		repo.updateId(obj, key)
// //
// // 		return key, err
// // }
// //
// func (repo *DataRepo) updateId(obj interface{}, key *datastore.Key) {
// 			switch x := obj.(type) {
// 			case *models.Ship:
// 				x.Id = key.Encode()
// 			}
// }
