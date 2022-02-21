package combination_test

import "github.com/asukakenji/go-benchmarks/combination"

func CombinationWithRepetition0[T any](a []T, r int) [][]T {
	if len(a) == 0 || r < 0 {
		return [][]T{}
	}
	if r == 0 {
		return [][]T{{}}
	}
	switch r {
	case 1:
		return CombinationWithRepetition0R1(a)
	case 2:
		return CombinationWithRepetition0R2(a)
	case 3:
		return CombinationWithRepetition0R3(a)
	case 4:
		return CombinationWithRepetition0R4(a)
	case 5:
		return CombinationWithRepetition0R5(a)
	case 6:
		return CombinationWithRepetition0R6(a)
	case 7:
		return CombinationWithRepetition0R7(a)
	case 8:
		return CombinationWithRepetition0R8(a)
	case 9:
		return CombinationWithRepetition0R9(a)
	case 10:
		return CombinationWithRepetition0R10(a)
	case 11:
		return CombinationWithRepetition0R11(a)
	case 12:
		return CombinationWithRepetition0R12(a)
	case 13:
		return CombinationWithRepetition0R13(a)
	case 14:
		return CombinationWithRepetition0R14(a)
	case 15:
		return CombinationWithRepetition0R15(a)
	case 16:
		return CombinationWithRepetition0R16(a)
	default:
		panic("not implemented")
	}
}

func CombinationWithRepetition0R1[T any](a []T) [][]T {
	n := len(a)
	count := combination.CombinationWithRepetitionCount(n, 1)
	result := make([][]T, count)
	i := 0
	for i1 := 0; i1 < n; i1++ {
		result[i] = []T{
			a[i1],
		}
		i++
	}
	return result
}

func CombinationWithRepetition0R2[T any](a []T) [][]T {
	n := len(a)
	count := combination.CombinationWithRepetitionCount(n, 2)
	result := make([][]T, count)
	i := 0
	for i1 := 0; i1 < n; i1++ {
		for i2 := i1; i2 < n; i2++ {
			result[i] = []T{
				a[i1], a[i2],
			}
			i++
		}
	}
	return result
}

func CombinationWithRepetition0R3[T any](a []T) [][]T {
	n := len(a)
	count := combination.CombinationWithRepetitionCount(n, 3)
	result := make([][]T, count)
	i := 0
	for i1 := 0; i1 < n; i1++ {
		for i2 := i1; i2 < n; i2++ {
			for i3 := i2; i3 < n; i3++ {
				result[i] = []T{
					a[i1], a[i2], a[i3],
				}
				i++
			}
		}
	}
	return result
}

func CombinationWithRepetition0R4[T any](a []T) [][]T {
	n := len(a)
	count := combination.CombinationWithRepetitionCount(n, 4)
	result := make([][]T, count)
	i := 0
	for i1 := 0; i1 < n; i1++ {
		for i2 := i1; i2 < n; i2++ {
			for i3 := i2; i3 < n; i3++ {
				for i4 := i3; i4 < n; i4++ {
					result[i] = []T{
						a[i1], a[i2], a[i3], a[i4],
					}
					i++
				}
			}
		}
	}
	return result
}

func CombinationWithRepetition0R5[T any](a []T) [][]T {
	n := len(a)
	count := combination.CombinationWithRepetitionCount(n, 5)
	result := make([][]T, count)
	i := 0
	for i1 := 0; i1 < n; i1++ {
		for i2 := i1; i2 < n; i2++ {
			for i3 := i2; i3 < n; i3++ {
				for i4 := i3; i4 < n; i4++ {
					for i5 := i4; i5 < n; i5++ {
						result[i] = []T{
							a[i1], a[i2], a[i3], a[i4], a[i5],
						}
						i++
					}
				}
			}
		}
	}
	return result
}

func CombinationWithRepetition0R6[T any](a []T) [][]T {
	n := len(a)
	count := combination.CombinationWithRepetitionCount(n, 6)
	result := make([][]T, count)
	i := 0
	for i1 := 0; i1 < n; i1++ {
		for i2 := i1; i2 < n; i2++ {
			for i3 := i2; i3 < n; i3++ {
				for i4 := i3; i4 < n; i4++ {
					for i5 := i4; i5 < n; i5++ {
						for i6 := i5; i6 < n; i6++ {
							result[i] = []T{
								a[i1], a[i2], a[i3], a[i4], a[i5], a[i6],
							}
							i++
						}
					}
				}
			}
		}
	}
	return result
}

func CombinationWithRepetition0R7[T any](a []T) [][]T {
	n := len(a)
	count := combination.CombinationWithRepetitionCount(n, 7)
	result := make([][]T, count)
	i := 0
	for i1 := 0; i1 < n; i1++ {
		for i2 := i1; i2 < n; i2++ {
			for i3 := i2; i3 < n; i3++ {
				for i4 := i3; i4 < n; i4++ {
					for i5 := i4; i5 < n; i5++ {
						for i6 := i5; i6 < n; i6++ {
							for i7 := i6; i7 < n; i7++ {
								result[i] = []T{
									a[i1], a[i2], a[i3], a[i4],
									a[i5], a[i6], a[i7],
								}
								i++
							}
						}
					}
				}
			}
		}
	}
	return result
}

func CombinationWithRepetition0R8[T any](a []T) [][]T {
	n := len(a)
	count := combination.CombinationWithRepetitionCount(n, 8)
	result := make([][]T, count)
	i := 0
	for i1 := 0; i1 < n; i1++ {
		for i2 := i1; i2 < n; i2++ {
			for i3 := i2; i3 < n; i3++ {
				for i4 := i3; i4 < n; i4++ {
					for i5 := i4; i5 < n; i5++ {
						for i6 := i5; i6 < n; i6++ {
							for i7 := i6; i7 < n; i7++ {
								for i8 := i7; i8 < n; i8++ {
									result[i] = []T{
										a[i1], a[i2], a[i3], a[i4],
										a[i5], a[i6], a[i7], a[i8],
									}
									i++
								}
							}
						}
					}
				}
			}
		}
	}
	return result
}

func CombinationWithRepetition0R9[T any](a []T) [][]T {
	n := len(a)
	count := combination.CombinationWithRepetitionCount(n, 9)
	result := make([][]T, count)
	i := 0
	for i1 := 0; i1 < n; i1++ {
		for i2 := i1; i2 < n; i2++ {
			for i3 := i2; i3 < n; i3++ {
				for i4 := i3; i4 < n; i4++ {
					for i5 := i4; i5 < n; i5++ {
						for i6 := i5; i6 < n; i6++ {
							for i7 := i6; i7 < n; i7++ {
								for i8 := i7; i8 < n; i8++ {
									for i9 := i8; i9 < n; i9++ {
										result[i] = []T{
											a[i1], a[i2], a[i3], a[i4],
											a[i5], a[i6], a[i7], a[i8],
											a[i9],
										}
										i++
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return result
}

func CombinationWithRepetition0R10[T any](a []T) [][]T {
	n := len(a)
	count := combination.CombinationWithRepetitionCount(n, 10)
	result := make([][]T, count)
	i := 0
	for i1 := 0; i1 < n; i1++ {
		for i2 := i1; i2 < n; i2++ {
			for i3 := i2; i3 < n; i3++ {
				for i4 := i3; i4 < n; i4++ {
					for i5 := i4; i5 < n; i5++ {
						for i6 := i5; i6 < n; i6++ {
							for i7 := i6; i7 < n; i7++ {
								for i8 := i7; i8 < n; i8++ {
									for i9 := i8; i9 < n; i9++ {
										for i10 := i9; i10 < n; i10++ {
											result[i] = []T{
												a[i1], a[i2], a[i3], a[i4],
												a[i5], a[i6], a[i7], a[i8],
												a[i9], a[i10],
											}
											i++
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return result
}

func CombinationWithRepetition0R11[T any](a []T) [][]T {
	n := len(a)
	count := combination.CombinationWithRepetitionCount(n, 11)
	result := make([][]T, count)
	i := 0
	for i1 := 0; i1 < n; i1++ {
		for i2 := i1; i2 < n; i2++ {
			for i3 := i2; i3 < n; i3++ {
				for i4 := i3; i4 < n; i4++ {
					for i5 := i4; i5 < n; i5++ {
						for i6 := i5; i6 < n; i6++ {
							for i7 := i6; i7 < n; i7++ {
								for i8 := i7; i8 < n; i8++ {
									for i9 := i8; i9 < n; i9++ {
										for i10 := i9; i10 < n; i10++ {
											for i11 := i10; i11 < n; i11++ {
												result[i] = []T{
													a[i1], a[i2], a[i3], a[i4],
													a[i5], a[i6], a[i7], a[i8],
													a[i9], a[i10], a[i11],
												}
												i++
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return result
}

func CombinationWithRepetition0R12[T any](a []T) [][]T {
	n := len(a)
	count := combination.CombinationWithRepetitionCount(n, 12)
	result := make([][]T, count)
	i := 0
	for i1 := 0; i1 < n; i1++ {
		for i2 := i1; i2 < n; i2++ {
			for i3 := i2; i3 < n; i3++ {
				for i4 := i3; i4 < n; i4++ {
					for i5 := i4; i5 < n; i5++ {
						for i6 := i5; i6 < n; i6++ {
							for i7 := i6; i7 < n; i7++ {
								for i8 := i7; i8 < n; i8++ {
									for i9 := i8; i9 < n; i9++ {
										for i10 := i9; i10 < n; i10++ {
											for i11 := i10; i11 < n; i11++ {
												for i12 := i11; i12 < n; i12++ {
													result[i] = []T{
														a[i1], a[i2], a[i3], a[i4],
														a[i5], a[i6], a[i7], a[i8],
														a[i9], a[i10], a[i11], a[i12],
													}
													i++
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return result
}

func CombinationWithRepetition0R13[T any](a []T) [][]T {
	n := len(a)
	count := combination.CombinationWithRepetitionCount(n, 13)
	result := make([][]T, count)
	i := 0
	for i1 := 0; i1 < n; i1++ {
		for i2 := i1; i2 < n; i2++ {
			for i3 := i2; i3 < n; i3++ {
				for i4 := i3; i4 < n; i4++ {
					for i5 := i4; i5 < n; i5++ {
						for i6 := i5; i6 < n; i6++ {
							for i7 := i6; i7 < n; i7++ {
								for i8 := i7; i8 < n; i8++ {
									for i9 := i8; i9 < n; i9++ {
										for i10 := i9; i10 < n; i10++ {
											for i11 := i10; i11 < n; i11++ {
												for i12 := i11; i12 < n; i12++ {
													for i13 := i12; i13 < n; i13++ {
														result[i] = []T{
															a[i1], a[i2], a[i3], a[i4],
															a[i5], a[i6], a[i7], a[i8],
															a[i9], a[i10], a[i11], a[i12],
															a[i13],
														}
														i++
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return result
}

func CombinationWithRepetition0R14[T any](a []T) [][]T {
	n := len(a)
	count := combination.CombinationWithRepetitionCount(n, 14)
	result := make([][]T, count)
	i := 0
	for i1 := 0; i1 < n; i1++ {
		for i2 := i1; i2 < n; i2++ {
			for i3 := i2; i3 < n; i3++ {
				for i4 := i3; i4 < n; i4++ {
					for i5 := i4; i5 < n; i5++ {
						for i6 := i5; i6 < n; i6++ {
							for i7 := i6; i7 < n; i7++ {
								for i8 := i7; i8 < n; i8++ {
									for i9 := i8; i9 < n; i9++ {
										for i10 := i9; i10 < n; i10++ {
											for i11 := i10; i11 < n; i11++ {
												for i12 := i11; i12 < n; i12++ {
													for i13 := i12; i13 < n; i13++ {
														for i14 := i13; i14 < n; i14++ {
															result[i] = []T{
																a[i1], a[i2], a[i3], a[i4],
																a[i5], a[i6], a[i7], a[i8],
																a[i9], a[i10], a[i11], a[i12],
																a[i13], a[i14],
															}
															i++
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return result
}

func CombinationWithRepetition0R15[T any](a []T) [][]T {
	n := len(a)
	count := combination.CombinationWithRepetitionCount(n, 15)
	result := make([][]T, count)
	i := 0
	for i1 := 0; i1 < n; i1++ {
		for i2 := i1; i2 < n; i2++ {
			for i3 := i2; i3 < n; i3++ {
				for i4 := i3; i4 < n; i4++ {
					for i5 := i4; i5 < n; i5++ {
						for i6 := i5; i6 < n; i6++ {
							for i7 := i6; i7 < n; i7++ {
								for i8 := i7; i8 < n; i8++ {
									for i9 := i8; i9 < n; i9++ {
										for i10 := i9; i10 < n; i10++ {
											for i11 := i10; i11 < n; i11++ {
												for i12 := i11; i12 < n; i12++ {
													for i13 := i12; i13 < n; i13++ {
														for i14 := i13; i14 < n; i14++ {
															for i15 := i14; i15 < n; i15++ {
																result[i] = []T{
																	a[i1], a[i2], a[i3], a[i4],
																	a[i5], a[i6], a[i7], a[i8],
																	a[i9], a[i10], a[i11], a[i12],
																	a[i13], a[i14], a[i15],
																}
																i++
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return result
}

func CombinationWithRepetition0R16[T any](a []T) [][]T {
	n := len(a)
	count := combination.CombinationWithRepetitionCount(n, 16)
	result := make([][]T, count)
	i := 0
	for i1 := 0; i1 < n; i1++ {
		for i2 := i1; i2 < n; i2++ {
			for i3 := i2; i3 < n; i3++ {
				for i4 := i3; i4 < n; i4++ {
					for i5 := i4; i5 < n; i5++ {
						for i6 := i5; i6 < n; i6++ {
							for i7 := i6; i7 < n; i7++ {
								for i8 := i7; i8 < n; i8++ {
									for i9 := i8; i9 < n; i9++ {
										for i10 := i9; i10 < n; i10++ {
											for i11 := i10; i11 < n; i11++ {
												for i12 := i11; i12 < n; i12++ {
													for i13 := i12; i13 < n; i13++ {
														for i14 := i13; i14 < n; i14++ {
															for i15 := i14; i15 < n; i15++ {
																for i16 := i15; i16 < n; i16++ {
																	result[i] = []T{
																		a[i1], a[i2], a[i3], a[i4],
																		a[i5], a[i6], a[i7], a[i8],
																		a[i9], a[i10], a[i11], a[i12],
																		a[i13], a[i14], a[i15], a[i16],
																	}
																	i++
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return result
}
