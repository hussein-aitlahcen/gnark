// Code generated by internal/gpoint DO NOT EDIT
package bn256

import (
	"fmt"
	"testing"

	"github.com/consensys/gnark/ecc/bn256/fr"
)

func TestG1JacToAffineFromJac(t *testing.T) {

	p := testPointsG1()

	_p := G1Affine{}
	p[0].ToAffineFromJac(&_p)
	if !_p.X.Equal(&p[1].X) || !_p.Y.Equal(&p[1].Y) {
		t.Fatal("ToAffineFromJac failed")
	}

}

func TestG1Conv(t *testing.T) {
	p := testPointsG1()

	for i := 0; i < len(p); i++ {
		var pJac G1Jac
		var pAff G1Affine
		p[i].ToAffineFromJac(&pAff)
		pAff.ToJacobian(&pJac)
		if !pJac.Equal(&p[i]) {
			t.Fatal("jacobian to affine to jacobian fails")
		}
	}
}

func TestG1JacAdd(t *testing.T) {

	curve := BN256()
	p := testPointsG1()

	// p3 = p1 + p2
	p1 := p[1].Clone()
	_p2 := G1Affine{}
	p[2].ToAffineFromJac(&_p2)
	p[1].AddMixed(&_p2)
	p[2].Add(curve, p1)

	if !p[3].Equal(&p[1]) {
		t.Fatal("Add failed")
	}

	// test commutativity
	if !p[3].Equal(&p[2]) {
		t.Fatal("Add failed")
	}
}

func TestG1JacSub(t *testing.T) {

	curve := BN256()
	p := testPointsG1()

	// p4 = p1 - p2
	p[1].Sub(curve, p[2])

	if !p[4].Equal(&p[1]) {
		t.Fatal("Sub failed")
	}
}

func TestG1JacDouble(t *testing.T) {

	curve := BN256()
	p := testPointsG1()

	// p5 = 2 * p1
	p[1].Double()
	if !p[5].Equal(&p[1]) {
		t.Fatal("Double failed")
	}

	G := curve.g1Infinity.Clone()
	R := curve.g1Infinity.Clone()
	G.Double()

	if !G.Equal(R) {
		t.Fatal("Double failed (infinity case)")
	}
}

func TestG1JacScalarMul(t *testing.T) {

	curve := BN256()
	p := testPointsG1()

	// p6 = [p1]32394 (scalar mul)
	scalar := fr.Element{32394}
	p[1].ScalarMul(curve, &p[1], scalar)

	if !p[1].Equal(&p[6]) {
		t.Error("ScalarMul failed")
	}
}

func TestG1JacMultiExp(t *testing.T) {
	curve := BN256()
	// var points []G1Jac
	var scalars []fr.Element
	var got G1Jac

	//
	// Test 1: testPointsG1multiExp
	//
	// TODO why is this commented?
	// numPoints, wants := testPointsG1MultiExpResults()

	// for i := range numPoints {
	// 	if numPoints[i] > 10000 {
	// 		continue
	// 	}
	// 	points, scalars = testPointsG1MultiExp(numPoints[i])

	// 	got.multiExp(curve, points, scalars)
	// 	if !got.Equal(&wants[i]) {
	// 		t.Error("multiExp G1Jac fail for points:", numPoints[i])
	// 	}
	// }

	//
	// Test 2: testPointsG1()
	//
	p := testPointsG1()

	// scalars
	s1 := fr.Element{23872983, 238203802, 9827897384, 2372}
	s2 := fr.Element{128923, 2878236, 398478, 187970707}
	s3 := fr.Element{9038947, 3947970, 29080823, 282739}

	scalars = []fr.Element{
		s1,
		s2,
		s3,
	}

	got.multiExp(curve, p[17:20], scalars)
	if !got.Equal(&p[20]) {
		t.Error("multiExp G1Jac failed")
	}

	//
	// Test 3: edge cases
	//

	// one input point p[1]
	scalars[0] = fr.Element{32394, 0, 0, 0} // single-word scalar
	got.multiExp(curve, p[1:2], scalars[:1])
	if !got.Equal(&p[6]) {
		t.Error("multiExp G1Jac failed, scalar:", scalars[0])
	}

	scalars[0] = fr.Element{2, 0, 0, 0} // scalar = 2
	got.multiExp(curve, p[1:2], scalars[:1])
	if !got.Equal(&p[5]) {
		t.Error("multiExp G1Jac failed, scalar:", scalars[0])
	}
	scalars[0] = fr.Element{1, 0, 0, 0} // scalar = 1
	got.multiExp(curve, p[1:2], scalars[:1])
	if !got.Equal(&p[1]) {
		t.Error("multiExp G1Jac failed, scalar:", scalars[0])
	}
	scalars[0] = fr.Element{0, 0, 0, 0} // scalar = 0
	got.multiExp(curve, p[1:2], scalars[:1])
	if !got.Equal(&curve.g1Infinity) {
		t.Error("multiExp G1Jac failed, scalar:", scalars[0])
	}
	scalars[0] = fr.Element{^uint64(0), ^uint64(0), ^uint64(0), ^uint64(0)} // scalar == (4-word maxuint)
	got.multiExp(curve, p[1:2], scalars[:1])
	if !got.Equal(&p[21]) {
		t.Error("multiExp G1Jac failed, scalar:", scalars[0])
	}

	// one input point curve.g1Infinity
	infinity := []G1Jac{curve.g1Infinity}

	scalars[0] = fr.Element{32394, 0, 0, 0} // single-word scalar
	got.multiExp(curve, infinity, scalars[:1])
	if !got.Equal(&curve.g1Infinity) {
		t.Error("multiExp G1Jac failed, scalar:", scalars[0])
	}
	scalars[0] = fr.Element{2, 0, 0, 0} // scalar = 2
	got.multiExp(curve, infinity, scalars[:1])
	if !got.Equal(&curve.g1Infinity) {
		t.Error("multiExp G1Jac failed, scalar:", scalars[0])
	}
	scalars[0] = fr.Element{1, 0, 0, 0} // scalar = 1
	got.multiExp(curve, infinity, scalars[:1])
	if !got.Equal(&curve.g1Infinity) {
		t.Error("multiExp G1Jac failed, scalar:", scalars[0])
	}
	scalars[0] = fr.Element{0, 0, 0, 0} // scalar = 0
	got.multiExp(curve, infinity, scalars[:1])
	if !got.Equal(&curve.g1Infinity) {
		t.Error("multiExp G1Jac failed, scalar:", scalars[0])
	}
	scalars[0] = fr.Element{^uint64(0), ^uint64(0), ^uint64(0), ^uint64(0)} // scalar == (4-word maxuint)
	got.multiExp(curve, infinity, scalars[:1])
	if !got.Equal(&curve.g1Infinity) {
		t.Error("multiExp G1Jac failed, scalar:", scalars[0])
	}

	// two input points: p[1], curve.g1Infinity
	twoPoints := []G1Jac{p[1], curve.g1Infinity}

	scalars[0] = fr.Element{32394, 0, 0, 0} // single-word scalar
	scalars[1] = fr.Element{2, 0, 0, 0}     // scalar = 2
	got.multiExp(curve, twoPoints, scalars[:2])
	if !got.Equal(&p[6]) {
		t.Error("multiExp G1Jac failed, scalar:", scalars[0])
	}
	scalars[0] = fr.Element{2, 0, 0, 0} // scalar = 2
	scalars[1] = fr.Element{1, 0, 0, 0} // scalar = 1
	got.multiExp(curve, twoPoints, scalars[:2])
	if !got.Equal(&p[5]) {
		t.Error("multiExp G1Jac failed, scalar:", scalars[0])
	}
	scalars[0] = fr.Element{1, 0, 0, 0} // scalar = 1
	scalars[1] = fr.Element{0, 0, 0, 0} // scalar = 0
	got.multiExp(curve, twoPoints, scalars[:2])
	if !got.Equal(&p[1]) {
		t.Error("multiExp G1Jac failed, scalar:", scalars[0])
	}
	scalars[0] = fr.Element{0, 0, 0, 0}                                     // scalar = 0
	scalars[1] = fr.Element{^uint64(0), ^uint64(0), ^uint64(0), ^uint64(0)} // scalar == (4-word maxuint)
	got.multiExp(curve, twoPoints, scalars[:2])
	if !got.Equal(&curve.g1Infinity) {
		t.Error("multiExp G1Jac failed, scalar:", scalars[0])
	}
	scalars[0] = fr.Element{^uint64(0), ^uint64(0), ^uint64(0), ^uint64(0)} // scalar == (4-word maxuint)
	scalars[1] = fr.Element{32394, 0, 0, 0}                                 // single-word scalar
	got.multiExp(curve, twoPoints, scalars[:2])
	if !got.Equal(&p[21]) {
		t.Error("multiExp G1Jac failed, scalar:", scalars[0])
	}

	// TODO: Jacobian points with nontrivial Z coord?
}

func TestMultiExpG1(t *testing.T) {

	curve := BN256()

	pointsJac := make([]G1Jac, 5)
	pointsAff := make([]G1Affine, 5)
	scalars := make([]fr.Element, 5)
	scalars[0].SetString("6833313093782752447774533032379533859360921141590695983").FromMont()
	scalars[1].SetString("6833313093782695347774533032379422124572402138975338593590695983").FromMont()
	scalars[2].SetString("683331309378269530181623992840859250215771777453309360695983").FromMont()
	scalars[3].SetString("6833353018162399284042212457240213897533859360921141590695983").FromMont()
	scalars[4].SetString("68333130937826953018162385525244777453303237942212485936983").FromMont()

	gens := make([]fr.Element, 5)
	gens[0].SetUint64(1).FromMont()
	gens[1].SetUint64(5).FromMont()
	gens[2].SetUint64(7).FromMont()
	gens[3].SetUint64(11).FromMont()
	gens[4].SetUint64(13).FromMont()

	pointsJac[0].ScalarMul(curve, &curve.g1Gen, gens[0])
	pointsJac[1].ScalarMul(curve, &curve.g1Gen, gens[1])
	pointsJac[2].ScalarMul(curve, &curve.g1Gen, gens[2])
	pointsJac[3].ScalarMul(curve, &curve.g1Gen, gens[3])
	pointsJac[4].ScalarMul(curve, &curve.g1Gen, gens[4])
	for i := 0; i < 5; i++ {
		pointsJac[i].ToAffineFromJac(&pointsAff[i])
	}

	pointsRes := make([]G1Jac, 5)
	pointsRes[0].ScalarMul(curve, &pointsJac[0], scalars[0])
	pointsRes[1].ScalarMul(curve, &pointsJac[1], scalars[1])
	pointsRes[2].ScalarMul(curve, &pointsJac[2], scalars[2])
	pointsRes[3].ScalarMul(curve, &pointsJac[3], scalars[3])
	pointsRes[4].ScalarMul(curve, &pointsJac[4], scalars[4])

	res := curve.g1Infinity

	for i := 0; i < 5; i++ {
		res.Add(curve, &pointsRes[i])
	}

	var multiExpRes G1Jac
	<-multiExpRes.MultiExp(curve, pointsAff, scalars)

	if !multiExpRes.Equal(&res) {
		fmt.Println("multiExp failed")
	}
}

func TestMultiExpG1LotOfPoints(t *testing.T) {

	curve := BN256()

	var G G1Jac

	samplePoints := make([]G1Affine, 1000)
	sampleScalars := make([]fr.Element, 1000)

	G.Set(&curve.g1Gen)

	for i := 1; i <= 1000; i++ {
		sampleScalars[i-1].SetUint64(uint64(i)).FromMont()
		G.ToAffineFromJac(&samplePoints[i-1])
	}

	var testPoint G1Jac

	<-testPoint.MultiExp(curve, samplePoints, sampleScalars)

	var finalScalar fr.Element
	finalScalar.SetUint64(500500).FromMont()
	var finalPoint G1Jac
	finalPoint.ScalarMul(curve, &G, finalScalar)

	if !finalPoint.Equal(&testPoint) {
		t.Fatal("error multi exp")
	}

}

func testPointsG1MultiExp(n int) (points []G1Jac, scalars []fr.Element) {

	curve := BN256()

	// points
	points = make([]G1Jac, n)
	points[0].Set(&curve.g1Gen)
	points[1].Set(&points[0]).Double() // can't call p.Add(a) when p equals a
	for i := 2; i < len(points); i++ {
		points[i].Set(&points[i-1]).Add(curve, &points[0]) // points[i] = i*g1Gen
	}

	// scalars
	// non-Montgomery form
	// cardinality of G1 is the fr modulus, so scalars should be fr.Elements
	// non-Montgomery form
	scalars = make([]fr.Element, n)

	// To ensure a diverse selection of scalars that use all words of an fr.Element,
	// each scalar should be a power of a large generator of fr.
	// 22 is a small generator of fr for bls377.
	// 2^{31}-1 is prime, so 22^{2^31}-1} is a large generator of fr for bls377
	// generator in Montgomery form
	var scalarGenMont fr.Element
	scalarGenMont.SetString("7716837800905789770901243404444209691916730933998574719964609384059111546487")

	scalars[0].Set(&scalarGenMont).FromMont()

	var curScalarMont fr.Element // Montgomery form
	curScalarMont.Set(&scalarGenMont)
	for i := 1; i < len(scalars); i++ {
		curScalarMont.MulAssign(&scalarGenMont)
		scalars[i].Set(&curScalarMont).FromMont() // scalars[i] = scalars[0]^i
	}

	return points, scalars
}

//--------------------//
//     benches		  //
//--------------------//

var benchResG1 G1Jac

func BenchmarkG1ScalarMul(b *testing.B) {

	curve := BN256()
	p := testPointsG1()

	var scalar fr.Element
	scalar.SetRandom()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p[1].ScalarMul(curve, &p[1], scalar)
		b.StopTimer()
		scalar.SetRandom()
		b.StartTimer()
	}

}

func BenchmarkG1Add(b *testing.B) {

	curve := BN256()
	p := testPointsG1()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchResG1 = p[1]
		benchResG1.Add(curve, &p[2])
	}

}

func BenchmarkG1AddMixed(b *testing.B) {

	p := testPointsG1()
	_p2 := G1Affine{}
	p[2].ToAffineFromJac(&_p2)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchResG1 = p[1]
		benchResG1.AddMixed(&_p2)
	}

}

func BenchmarkG1Double(b *testing.B) {

	p := testPointsG1()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchResG1 = p[1]
		benchResG1.Double()
	}

}

func BenchmarkG1WindowedMultiExp(b *testing.B) {
	curve := BN256()

	var G G1Jac

	var mixer fr.Element
	mixer.SetString("7716837800905789770901243404444209691916730933998574719964609384059111546487")

	var nbSamples int
	nbSamples = 400000

	samplePoints := make([]G1Jac, nbSamples)
	sampleScalars := make([]fr.Element, nbSamples)

	G.Set(&curve.g1Gen)

	for i := 1; i <= nbSamples; i++ {
		sampleScalars[i-1].SetUint64(uint64(i)).
			Mul(&sampleScalars[i-1], &mixer).
			FromMont()
		samplePoints[i-1].Set(&curve.g1Gen)
	}

	var testPoint G1Jac

	for i := 0; i < 8; i++ {
		b.Run(fmt.Sprintf("%d points", (i+1)*50000), func(b *testing.B) {
			b.ResetTimer()
			for j := 0; j < b.N; j++ {
				testPoint.WindowedMultiExp(curve, samplePoints[:50000+i*50000], sampleScalars[:50000+i*50000])
			}
		})
	}
}

func BenchmarkMultiExpG1(b *testing.B) {

	curve := BN256()

	var G G1Jac

	var mixer fr.Element
	mixer.SetString("7716837800905789770901243404444209691916730933998574719964609384059111546487")

	var nbSamples int
	nbSamples = 400000

	samplePoints := make([]G1Affine, nbSamples)
	sampleScalars := make([]fr.Element, nbSamples)

	G.Set(&curve.g1Gen)

	for i := 1; i <= nbSamples; i++ {
		sampleScalars[i-1].SetUint64(uint64(i)).
			Mul(&sampleScalars[i-1], &mixer).
			FromMont()
		G.ToAffineFromJac(&samplePoints[i-1])
	}

	var testPoint G1Jac

	for i := 0; i < 8; i++ {
		b.Run(fmt.Sprintf("%d points", (i+1)*50000), func(b *testing.B) {
			b.ResetTimer()
			for j := 0; j < b.N; j++ {
				<-testPoint.MultiExp(curve, samplePoints[:50000+i*50000], sampleScalars[:50000+i*50000])
			}
		})
	}

}
