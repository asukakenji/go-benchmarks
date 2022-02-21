package combination_test

import "github.com/asukakenji/go-benchmarks/combination"

func Combination0[T any](a []T, r int) [][]T {
	if r < 0 || r > len(a) {
		return [][]T{}
	}
	if r == 0 {
		return [][]T{{}}
	}
	switch r {
	case 1:
		return Combination0R1(a)
	case 2:
		return Combination0R2(a)
	case 3:
		return Combination0R3(a)
	case 4:
		return Combination0R4(a)
	case 5:
		return Combination0R5(a)
	case 6:
		return Combination0R6(a)
	case 7:
		return Combination0R7(a)
	case 8:
		return Combination0R8(a)
	case 9:
		return Combination0R9(a)
	case 10:
		return Combination0R10(a)
	case 11:
		return Combination0R11(a)
	case 12:
		return Combination0R12(a)
	case 13:
		return Combination0R13(a)
	case 14:
		return Combination0R14(a)
	case 15:
		return Combination0R15(a)
	case 16:
		return Combination0R16(a)
	default:
		panic("not implemented")
	}
}

func Combination0R1[T any](a []T) [][]T {
	n := len(a)
	count := combination.CombinationCount(n, 1)
	result := make([][]T, 0, count)
	for i1 := 0; i1 < n; i1++ {
		result = append(result, []T{
			a[i1],
		})
	}
	return result
}

func Combination0R2[T any](a []T) [][]T {
	n := len(a)
	count := combination.CombinationCount(n, 2)
	result := make([][]T, 0, count)
	for i1 := 0; i1 < n-1; i1++ {
		for i2 := i1 + 1; i2 < n; i2++ {
			result = append(result, []T{
				a[i1], a[i2],
			})
		}
	}
	return result
}

func Combination0R3[T any](a []T) [][]T {
	n := len(a)
	count := combination.CombinationCount(n, 3)
	result := make([][]T, 0, count)
	for i1 := 0; i1 < n-2; i1++ {
		for i2 := i1 + 1; i2 < n-1; i2++ {
			for i3 := i2 + 1; i3 < n; i3++ {
				result = append(result, []T{
					a[i1], a[i2], a[i3],
				})
			}
		}
	}
	return result
}

func Combination0R4[T any](a []T) [][]T {
	n := len(a)
	count := combination.CombinationCount(n, 4)
	result := make([][]T, 0, count)
	for i1 := 0; i1 < n-3; i1++ {
		for i2 := i1 + 1; i2 < n-2; i2++ {
			for i3 := i2 + 1; i3 < n-1; i3++ {
				for i4 := i3 + 1; i4 < n; i4++ {
					result = append(result, []T{
						a[i1], a[i2], a[i3], a[i4],
					})
				}
			}
		}
	}
	return result
}

func Combination0R5[T any](a []T) [][]T {
	n := len(a)
	count := combination.CombinationCount(n, 5)
	result := make([][]T, 0, count)
	for i1 := 0; i1 < n-4; i1++ {
		for i2 := i1 + 1; i2 < n-3; i2++ {
			for i3 := i2 + 1; i3 < n-2; i3++ {
				for i4 := i3 + 1; i4 < n-1; i4++ {
					for i5 := i4 + 1; i5 < n; i5++ {
						result = append(result, []T{
							a[i1], a[i2], a[i3], a[i4], a[i5],
						})
					}
				}
			}
		}
	}
	return result
}

func Combination0R6[T any](a []T) [][]T {
	n := len(a)
	count := combination.CombinationCount(n, 6)
	result := make([][]T, 0, count)
	for i1 := 0; i1 < n-5; i1++ {
		for i2 := i1 + 1; i2 < n-4; i2++ {
			for i3 := i2 + 1; i3 < n-3; i3++ {
				for i4 := i3 + 1; i4 < n-2; i4++ {
					for i5 := i4 + 1; i5 < n-1; i5++ {
						for i6 := i5 + 1; i6 < n; i6++ {
							result = append(result, []T{
								a[i1], a[i2], a[i3], a[i4], a[i5], a[i6],
							})
						}
					}
				}
			}
		}
	}
	return result
}

func Combination0R7[T any](a []T) [][]T {
	n := len(a)
	count := combination.CombinationCount(n, 7)
	result := make([][]T, 0, count)
	for i1 := 0; i1 < n-6; i1++ {
		for i2 := i1 + 1; i2 < n-5; i2++ {
			for i3 := i2 + 1; i3 < n-4; i3++ {
				for i4 := i3 + 1; i4 < n-3; i4++ {
					for i5 := i4 + 1; i5 < n-2; i5++ {
						for i6 := i5 + 1; i6 < n-1; i6++ {
							for i7 := i6 + 1; i7 < n; i7++ {
								result = append(result, []T{
									a[i1], a[i2], a[i3], a[i4],
									a[i5], a[i6], a[i7],
								})
							}
						}
					}
				}
			}
		}
	}
	return result
}

func Combination0R8[T any](a []T) [][]T {
	n := len(a)
	count := combination.CombinationCount(n, 8)
	result := make([][]T, 0, count)
	for i1 := 0; i1 < n-7; i1++ {
		for i2 := i1 + 1; i2 < n-6; i2++ {
			for i3 := i2 + 1; i3 < n-5; i3++ {
				for i4 := i3 + 1; i4 < n-4; i4++ {
					for i5 := i4 + 1; i5 < n-3; i5++ {
						for i6 := i5 + 1; i6 < n-2; i6++ {
							for i7 := i6 + 1; i7 < n-1; i7++ {
								for i8 := i7 + 1; i8 < n; i8++ {
									result = append(result, []T{
										a[i1], a[i2], a[i3], a[i4],
										a[i5], a[i6], a[i7], a[i8],
									})
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

func Combination0R9[T any](a []T) [][]T {
	n := len(a)
	count := combination.CombinationCount(n, 9)
	result := make([][]T, 0, count)
	for i1 := 0; i1 < n-8; i1++ {
		for i2 := i1 + 1; i2 < n-7; i2++ {
			for i3 := i2 + 1; i3 < n-6; i3++ {
				for i4 := i3 + 1; i4 < n-5; i4++ {
					for i5 := i4 + 1; i5 < n-4; i5++ {
						for i6 := i5 + 1; i6 < n-3; i6++ {
							for i7 := i6 + 1; i7 < n-2; i7++ {
								for i8 := i7 + 1; i8 < n-1; i8++ {
									for i9 := i8 + 1; i9 < n; i9++ {
										result = append(result, []T{
											a[i1], a[i2], a[i3], a[i4],
											a[i5], a[i6], a[i7], a[i8],
											a[i9],
										})
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

func Combination0R10[T any](a []T) [][]T {
	n := len(a)
	count := combination.CombinationCount(n, 10)
	result := make([][]T, 0, count)
	for i1 := 0; i1 < n-9; i1++ {
		for i2 := i1 + 1; i2 < n-8; i2++ {
			for i3 := i2 + 1; i3 < n-7; i3++ {
				for i4 := i3 + 1; i4 < n-6; i4++ {
					for i5 := i4 + 1; i5 < n-5; i5++ {
						for i6 := i5 + 1; i6 < n-4; i6++ {
							for i7 := i6 + 1; i7 < n-3; i7++ {
								for i8 := i7 + 1; i8 < n-2; i8++ {
									for i9 := i8 + 1; i9 < n-1; i9++ {
										for i10 := i9 + 1; i10 < n; i10++ {
											result = append(result, []T{
												a[i1], a[i2], a[i3], a[i4],
												a[i5], a[i6], a[i7], a[i8],
												a[i9], a[i10],
											})
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

func Combination0R11[T any](a []T) [][]T {
	n := len(a)
	count := combination.CombinationCount(n, 10)
	result := make([][]T, 0, count)
	for i1 := 0; i1 < n-10; i1++ {
		for i2 := i1 + 1; i2 < n-9; i2++ {
			for i3 := i2 + 1; i3 < n-8; i3++ {
				for i4 := i3 + 1; i4 < n-7; i4++ {
					for i5 := i4 + 1; i5 < n-6; i5++ {
						for i6 := i5 + 1; i6 < n-5; i6++ {
							for i7 := i6 + 1; i7 < n-4; i7++ {
								for i8 := i7 + 1; i8 < n-3; i8++ {
									for i9 := i8 + 1; i9 < n-2; i9++ {
										for i10 := i9 + 1; i10 < n-1; i10++ {
											for i11 := i10 + 1; i11 < n; i11++ {
												result = append(result, []T{
													a[i1], a[i2], a[i3], a[i4],
													a[i5], a[i6], a[i7], a[i8],
													a[i9], a[i10], a[i11],
												})
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

func Combination0R12[T any](a []T) [][]T {
	n := len(a)
	count := combination.CombinationCount(n, 10)
	result := make([][]T, 0, count)
	for i1 := 0; i1 < n-11; i1++ {
		for i2 := i1 + 1; i2 < n-10; i2++ {
			for i3 := i2 + 1; i3 < n-9; i3++ {
				for i4 := i3 + 1; i4 < n-8; i4++ {
					for i5 := i4 + 1; i5 < n-7; i5++ {
						for i6 := i5 + 1; i6 < n-6; i6++ {
							for i7 := i6 + 1; i7 < n-5; i7++ {
								for i8 := i7 + 1; i8 < n-4; i8++ {
									for i9 := i8 + 1; i9 < n-3; i9++ {
										for i10 := i9 + 1; i10 < n-2; i10++ {
											for i11 := i10 + 1; i11 < n-1; i11++ {
												for i12 := i11 + 1; i12 < n; i12++ {
													result = append(result, []T{
														a[i1], a[i2], a[i3], a[i4],
														a[i5], a[i6], a[i7], a[i8],
														a[i9], a[i10], a[i11], a[i12],
													})
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

func Combination0R13[T any](a []T) [][]T {
	n := len(a)
	count := combination.CombinationCount(n, 10)
	result := make([][]T, 0, count)
	for i1 := 0; i1 < n-12; i1++ {
		for i2 := i1 + 1; i2 < n-11; i2++ {
			for i3 := i2 + 1; i3 < n-10; i3++ {
				for i4 := i3 + 1; i4 < n-9; i4++ {
					for i5 := i4 + 1; i5 < n-8; i5++ {
						for i6 := i5 + 1; i6 < n-7; i6++ {
							for i7 := i6 + 1; i7 < n-6; i7++ {
								for i8 := i7 + 1; i8 < n-5; i8++ {
									for i9 := i8 + 1; i9 < n-4; i9++ {
										for i10 := i9 + 1; i10 < n-3; i10++ {
											for i11 := i10 + 1; i11 < n-2; i11++ {
												for i12 := i11 + 1; i12 < n-1; i12++ {
													for i13 := i12 + 1; i13 < n; i13++ {
														result = append(result, []T{
															a[i1], a[i2], a[i3], a[i4],
															a[i5], a[i6], a[i7], a[i8],
															a[i9], a[i10], a[i11], a[i12],
															a[i13],
														})
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

func Combination0R14[T any](a []T) [][]T {
	n := len(a)
	count := combination.CombinationCount(n, 10)
	result := make([][]T, 0, count)
	for i1 := 0; i1 < n-13; i1++ {
		for i2 := i1 + 1; i2 < n-12; i2++ {
			for i3 := i2 + 1; i3 < n-11; i3++ {
				for i4 := i3 + 1; i4 < n-10; i4++ {
					for i5 := i4 + 1; i5 < n-9; i5++ {
						for i6 := i5 + 1; i6 < n-8; i6++ {
							for i7 := i6 + 1; i7 < n-7; i7++ {
								for i8 := i7 + 1; i8 < n-6; i8++ {
									for i9 := i8 + 1; i9 < n-5; i9++ {
										for i10 := i9 + 1; i10 < n-4; i10++ {
											for i11 := i10 + 1; i11 < n-3; i11++ {
												for i12 := i11 + 1; i12 < n-2; i12++ {
													for i13 := i12 + 1; i13 < n-1; i13++ {
														for i14 := i13 + 1; i14 < n; i14++ {
															result = append(result, []T{
																a[i1], a[i2], a[i3], a[i4],
																a[i5], a[i6], a[i7], a[i8],
																a[i9], a[i10], a[i11], a[i12],
																a[i13], a[i14],
															})
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

func Combination0R15[T any](a []T) [][]T {
	n := len(a)
	count := combination.CombinationCount(n, 10)
	result := make([][]T, 0, count)
	for i1 := 0; i1 < n-14; i1++ {
		for i2 := i1 + 1; i2 < n-13; i2++ {
			for i3 := i2 + 1; i3 < n-12; i3++ {
				for i4 := i3 + 1; i4 < n-11; i4++ {
					for i5 := i4 + 1; i5 < n-10; i5++ {
						for i6 := i5 + 1; i6 < n-9; i6++ {
							for i7 := i6 + 1; i7 < n-8; i7++ {
								for i8 := i7 + 1; i8 < n-7; i8++ {
									for i9 := i8 + 1; i9 < n-6; i9++ {
										for i10 := i9 + 1; i10 < n-5; i10++ {
											for i11 := i10 + 1; i11 < n-4; i11++ {
												for i12 := i11 + 1; i12 < n-3; i12++ {
													for i13 := i12 + 1; i13 < n-2; i13++ {
														for i14 := i13 + 1; i14 < n-1; i14++ {
															for i15 := i14 + 1; i15 < n; i15++ {
																result = append(result, []T{
																	a[i1], a[i2], a[i3], a[i4],
																	a[i5], a[i6], a[i7], a[i8],
																	a[i9], a[i10], a[i11], a[i12],
																	a[i13], a[i14], a[i15],
																})
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

func Combination0R16[T any](a []T) [][]T {
	n := len(a)
	count := combination.CombinationCount(n, 10)
	result := make([][]T, 0, count)
	for i1 := 0; i1 < n-15; i1++ {
		for i2 := i1 + 1; i2 < n-14; i2++ {
			for i3 := i2 + 1; i3 < n-13; i3++ {
				for i4 := i3 + 1; i4 < n-12; i4++ {
					for i5 := i4 + 1; i5 < n-11; i5++ {
						for i6 := i5 + 1; i6 < n-10; i6++ {
							for i7 := i6 + 1; i7 < n-9; i7++ {
								for i8 := i7 + 1; i8 < n-8; i8++ {
									for i9 := i8 + 1; i9 < n-7; i9++ {
										for i10 := i9 + 1; i10 < n-6; i10++ {
											for i11 := i10 + 1; i11 < n-5; i11++ {
												for i12 := i11 + 1; i12 < n-4; i12++ {
													for i13 := i12 + 1; i13 < n-3; i13++ {
														for i14 := i13 + 1; i14 < n-2; i14++ {
															for i15 := i14 + 1; i15 < n-1; i15++ {
																for i16 := i15 + 1; i16 < n; i16++ {
																	result = append(result, []T{
																		a[i1], a[i2], a[i3], a[i4],
																		a[i5], a[i6], a[i7], a[i8],
																		a[i9], a[i10], a[i11], a[i12],
																		a[i13], a[i14], a[i15], a[i16],
																	})
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
